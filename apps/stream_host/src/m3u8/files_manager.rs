use std::time::Instant;

use tokio::fs::{self, File};
use tokio::io::AsyncWriteExt;

use super::transport_stream::{SegmentMetadata, utils};
use super::{M3u8Config, M3u8Error};

const PLAYLIST_FILE: &str = "playlist.m3u8";

#[derive(Clone)]
pub struct M3u8FilesManager {
    pub(super) config: M3u8Config,
    pub(super) segment_metadata: SegmentMetadata,
}

impl M3u8FilesManager {
    pub fn new(config: M3u8Config, stream_id: String) -> Self {
        Self {
            segment_metadata: SegmentMetadata::new(config.streams_directory.join(stream_id)),
            config,
        }
    }

    pub fn should_try_next_segment(&self, timestamp: Instant) -> bool {
        timestamp.duration_since(self.segment_metadata.last_write) >= self.config.segment_duration
    }

    pub async fn initialize(&self) -> Result<(), M3u8Error> {
        fs::create_dir_all(&self.segment_metadata.directory)
            .await
            .map_err(|e| M3u8Error::InitializationError(e))
    }

    pub async fn clear(self) -> Result<(), M3u8Error> {
        fs::remove_dir_all(self.segment_metadata.directory)
            .await
            .map_err(|e| M3u8Error::ClearError(e))
    }

    pub(super) async fn clear_unused(&self) -> Result<(), M3u8Error> {
        if let Some(n) = self.config.clear_unused_segments {
            let idx = (self.segment_metadata.segment_index as isize)
                - (self.config.segments_per_stream + n) as isize;

            if idx >= 0 {
                utils::remove_segment(self.segment_metadata.directory.clone(), idx as usize)
                    .await
                    .map_err(|e| M3u8Error::ClearError(e))?;
            }
        }

        Ok(())
    }

    pub(super) async fn update_playlist(&self) -> Result<(), M3u8Error> {
        let mut playlist = File::options()
            .truncate(true)
            .create(true)
            .write(true)
            .open(self.segment_metadata.directory.join(PLAYLIST_FILE))
            .await
            .map_err(|e| M3u8Error::OpenPlaylistFileError(e))?;

        let mut buf = vec![];

        // * Write .m3u8 header.
        // ? [
        // ?     "#EXTM3U",
        // ?     "#EXT-X-VERSION:3",
        // ?     "#EXT-X-TARGETDURATION:<Duration in seconds>",
        // ?     "#EXT-X-MEDIA-SEQUENCE:<last_segment_index>",
        // ? ]
        buf.extend_from_slice(
            format!(
                "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-TARGETDURATION:{}\n#EXT-X-MEDIA-SEQUENCE:{}\n",
                self.config.segment_duration.as_millis().div_ceil(1000),
                self.segment_metadata.segment_index.saturating_sub(self.config.segments_per_stream)
            )
            .as_bytes(),
        );

        for idx in
            self.segment_metadata.segment_index.saturating_sub(self.config.segments_per_stream)
                ..self.segment_metadata.segment_index
        {
            buf.extend_from_slice(
                format!(
                    "#EXTINF:{:.3},\n{}\n",
                    self.config.segment_duration.as_secs_f32(),
                    utils::get_filename(idx)
                )
                .as_bytes(),
            )
        }

        playlist.write_all(&buf).await.map_err(|_| M3u8Error::WritePlaylistFileError)?;

        Ok(())
    }
}
