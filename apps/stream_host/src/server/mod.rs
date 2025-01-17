use std::path::{Path, PathBuf};
use std::sync::Arc;
use std::time::{Duration, Instant};

use futures::{StreamExt, TryFutureExt, TryStreamExt};
use log::{error, info};
use srt_protocol::settings::KeySettings;
use srt_tokio::{ConnectionRequest, SrtListener, SrtSocket};
use tokio::fs::File;
use tokio::io::{self, AsyncWriteExt};
use tokio::net::{TcpListener, TcpStream};
use tokio::sync::Mutex;
use tokio::task::JoinHandle;
use tokio::{fs, time};

// ? Maybe useful: "#EXT-X-DISCONTINUITY" @gpt "Use if switching codecs, resolutions, or timestamp discontinuities."
// const BUFFER_SIZE: usize = ?;
const SEGMENT_DURATION: Duration = Duration::from_millis(1_000);
const SEGMENTS_PER_STREAM: usize = 5;

#[derive(Debug)]
pub enum ServerError {
    FailedToAcceptSrtConnection(io::Error),
    FailedToCreate(io::Error),
    FailedToWrite(io::Error),
    ConnectionClosed,
    StreamIdNotProvided,
}

pub struct Server {
    content_path: &'static Path,
}

impl Server {
    pub fn new() -> Self {
        Self {
            content_path: Path::new("content/streams"),
        }
    }

    pub async fn start(self, port: u16) -> Result<(), Box<dyn std::error::Error>> {
        let (_server, mut incoming) = SrtListener::builder().bind(port).await?;

        let shared = Arc::new(self);
        while let Some(connection_req) = incoming.incoming().next().await {
            let shared_clone = shared.clone();
            tokio::spawn(async move {
                if let Err(err) = shared_clone.handle_srt_connection(connection_req).await {
                    error!("Error: {:?}", err);
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
                    if timestamp.duration_since(last_write_time) >= SEGMENT_DURATION {
                        // Flush previous file
                        segment_file.flush().await.map_err(|e| ServerError::FailedToWrite(e))?;

                        // Create new file
                        last_write_time = timestamp;
                        segment_index += 1;
                        segment_file = create_segment_file(segment_index).await?;

                        if segment_index % SEGMENTS_PER_STREAM == 0 {
                            tokio::spawn(update_m3u8_playlist(content_path.clone(), segment_index));
                        }
                    }

                    segment_file
                        .write_all(&data)
                        .await
                        .map_err(|e| ServerError::FailedToWrite(e))?;
                },
                Err(err) => {
                    error!("Error: {}", err);
                    return Err(ServerError::ConnectionClosed);
                    // info!("Connection closed: {:?}", v);
                },
                Ok(None) => {
                    error!("Shouldn't be possible");
                    // TODO: clear folder after stream ned
                    return Err(ServerError::ConnectionClosed);
                },
            }
        }

        Ok(())
    }
}

// TODO: Result<(), SomeError>
async fn update_m3u8_playlist(
    content_path: PathBuf,
    last_segment_index: usize,
) -> Result<(), String> {
    let mut playlist = File::options()
        .truncate(true)
        .create(true)
        .read(true)
        .write(true)
        .open(content_path.join("playlist.m3u8"))
        .await
        .map_err(|e| e.to_string())?;

    if last_segment_index < SEGMENTS_PER_STREAM-1 {
        return Ok(());
    }
    // TODO: Clear previous files

    playlist
        .write(
            // Same as
            // [
            //     "#EXTM3U",
            //     "#EXT-X-VERSION:3",
            //     "#EXT-X-TARGETDURATION:3",
            //     &format!("#EXT-X-MEDIA-SEQUENCE:{}", last_segment_index),
            // ]
            // .join("\n")
            format!(
                "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-TARGETDURATION:{duration}\n#EXT-X-MEDIA-SEQUENCE:{first_segment_index}\n",
                duration=SEGMENT_DURATION.as_millis().div_ceil(1000),
                first_segment_index=last_segment_index-SEGMENTS_PER_STREAM
            )
            .as_bytes(),
        )
        .await
        .map_err(|e| e.to_string())?;

    for i in last_segment_index - SEGMENTS_PER_STREAM..=last_segment_index {
        if i == 0 {

        }

        playlist
            .write(
                format!(
                    "#EXTINF:{duration:.3},\n{filename}\n",
                    duration=SEGMENT_DURATION.as_secs_f32(),
                    filename=format!("segment_{i}.ts")
                )
                .as_bytes(),
            )
            .await
            .map_err(|e| e.to_string())?;
    }

    Ok(())
}
