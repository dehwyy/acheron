use std::path::{Path, PathBuf};
use std::sync::Arc;
use std::time::{Duration, Instant};

use futures::{StreamExt, TryFutureExt, TryStreamExt, future};
use log::{error, info};
use srt_protocol::settings::KeySettings;
use srt_tokio::{ConnectionRequest, SrtListener, SrtSocket};
use tokio::fs::File;
use tokio::io::{self, AsyncWriteExt};
use tokio::net::{TcpListener, TcpStream};
use tokio::sync::Mutex;
use tokio::task::JoinHandle;
use tokio::{fs, time};

// ? Maybe useful: "#EXT-X-DISCONTINUITY" @gpt "Use if switching codecs, resolutions, or timestamp
// discontinuities." const BUFFER_SIZE: usize = ?;
// const segment_duration: Duration = Duration::from_millis(2000);
// const segments_per_stream: usize = 13;

#[derive(Debug)]
pub enum ServerError {
    FailedToAcceptSrtConnection(io::Error),
    FailedToDelete,
    FailedToCreate(io::Error),
    FailedToWrite(io::Error),
    ConnectionClosed,
    StreamIdNotProvided,
}

pub struct Server {
    segment_duration: Duration,
    segments_per_stream: usize,
    content_path: &'static Path,
}

impl Server {
    pub fn new(segment_duration_ms: u64, segments_per_stream: usize) -> Self {
        Self {
            segment_duration: Duration::from_millis(segment_duration_ms),
            segments_per_stream,
            content_path: Path::new("content/streams"),
        }
    }

    pub async fn start(self, port: u16) -> Result<(), Box<dyn std::error::Error>> {
        let (_server, mut incoming) = SrtListener::builder().bind(port).await?;

        info!("Server: SegmentDuration {dur}, SegmentsPerStream {segments}", dur = self.segment_duration.as_millis(), segments = self.segments_per_stream);
        info!("Listening on port {}", port);

        let shared = Arc::new(self);
        while let Some(connection_req) = incoming.incoming().next().await {
            let shared_clone = shared.clone();
            tokio::spawn(async move {
                if let Err(err) = shared_clone.handle_srt_connection(connection_req).await {
                    error!("{:?}", err);
                }
            });
        }

        Ok(())
    }

    async fn handle_srt_connection(
        self: Arc<Self>,
        conn: ConnectionRequest,
    ) -> Result<(), ServerError> {
        // Todo: perform some validation / authentication.
        let key_settings: Option<KeySettings> = None;

        let remote_addr = conn.remote();
        let stream_id = conn.stream_id().ok_or(ServerError::StreamIdNotProvided)?.to_string();

        info!("New connection: {remote_addr}");
        let mut socket = conn
            .accept(key_settings)
            .await
            .map_err(|e| ServerError::FailedToAcceptSrtConnection(e))?;

        let content_path = self.content_path.join(stream_id);
        // fs::remove_dir_all(&content_path).await.map_err(|e| ServerError::FailedToDelete)?;

        fs::create_dir_all(content_path.clone())
            .await
            .map_err(|e| ServerError::FailedToCreate(e))?;

        let create_filename =
            |segment_index: usize| content_path.join(format!("segment_{segment_index}.ts"));
        let create_segment_file = |segmnet_index: usize| {
            File::create(create_filename(segmnet_index)).map_err(|e| ServerError::FailedToCreate(e))
        };

        let mut segment_index = 0usize;
        let mut segment_file = create_segment_file(segment_index).await?;
        let mut last_write_time = Instant::now();

        loop {
            match socket.try_next().await {
                Ok(Some((timestamp, data))) => {
                    info!("Got data: {} bytes", data.len());
                    if timestamp.duration_since(last_write_time) >= self.segment_duration {
                        // Flush previous file
                        segment_file.flush().await.map_err(|e| ServerError::FailedToWrite(e))?;

                        // Create new file
                        last_write_time = timestamp;
                        segment_index += 1;
                        segment_file = create_segment_file(segment_index).await?;

                        tokio::spawn(update_m3u8_playlist(self.segment_duration, self.segments_per_stream, content_path.clone(), segment_index - 1));
                    }

                    segment_file
                        .write_all(&data)
                        .await
                        .map_err(|e| ServerError::FailedToWrite(e))?;
                },
                Err(err) => {
                    error!("WASSSS");

                    return Err(ServerError::ConnectionClosed);
                },
                Ok(None) => {
                    info!("Connection closed");
                    fs::remove_dir_all(content_path).await;
                    return Err(ServerError::ConnectionClosed);
                },
            }
        }

        Ok(())
    }
}

// TODO: Result<(), SomeError>
async fn update_m3u8_playlist(
    segment_duration: Duration,
    segments_per_stream: usize,
    content_path: PathBuf,
    last_segment_index: usize,
) -> Result<(), String> {

    // * Clear unused file
    if last_segment_index >= segments_per_stream + 5 {
        tokio::spawn(fs::remove_file(
            content_path
                .clone()
                .join(format!("segment_{idx}.ts", idx = last_segment_index - segments_per_stream - 5 )),
        ));
    }

    let mut playlist = File::options()
        .truncate(true)
        .create(true)
        .read(true)
        .write(true)
        .open(content_path.join("playlist.m3u8"))
        .await
        .map_err(|e| e.to_string())?;

    // * Write .m3u8 header.
    // ? [
    // ?     "#EXTM3U",
    // ?     "#EXT-X-VERSION:3",
    // ?     "#EXT-X-TARGETDURATION:<Duration in seconds>",
    // ?     "#EXT-X-MEDIA-SEQUENCE:<last_segment_index>",
    // ? ]

    let mut buf = vec![];
    buf.extend_from_slice(
            format!(
                "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-TARGETDURATION:{duration}\n#EXT-X-MEDIA-SEQUENCE:{first_segment_index}\n",
                duration=segment_duration.as_millis().div_ceil(1000),
                first_segment_index=last_segment_index.saturating_sub(segments_per_stream)
            )
            .as_bytes(),
    );

    for idx in last_segment_index.saturating_sub(segments_per_stream)..=last_segment_index {
        buf.extend_from_slice(
            format!(
                "#EXTINF:{duration:.3},\n{filename}\n",
                duration = segment_duration.as_secs_f32(),
                filename = format!("segment_{idx}.ts")
            )
            .as_bytes(),
        )
    }

    playlist.write_all(&buf).await.map_err(|e| e.to_string())?;

    Ok(())
}
