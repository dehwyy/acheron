use std::time::Instant;


mod transport_stream;
use log::error;
use transport_stream::SegmentFile;

mod errors;
pub use errors::*;

mod config;
pub use config::M3u8Config;

mod files_manager;
use files_manager::M3u8FilesManager;

pub struct M3u8 {
    manager: M3u8FilesManager,
    segment: SegmentFile,
}

impl M3u8 {
    /// Creates new instance of M3u8 + initializes dependecies.
    pub async fn try_new(config: M3u8Config, stream_id: String) -> Result<Self, M3u8Error> {
        let manager = M3u8FilesManager::new(config, stream_id);
        manager.initialize().await?;

        Ok(Self {
            segment: SegmentFile::try_new(manager.segment_metadata.directory.clone())
                .await
                .map_err(|e| M3u8Error::SegmentError(e))?,
            manager,
        })
    }

    pub async fn handle(&mut self, timestamp: Instant, data: &[u8]) -> Result<(), M3u8Error> {
        if self.manager.should_try_next_segment(timestamp) {
            self.segment
                .try_next(&mut self.manager.segment_metadata.segment_index)
                .await
                .map_err(|e| M3u8Error::SegmentError(e))?;

            self.manager.segment_metadata.last_write = timestamp;

            let manager = self.manager.clone();
            tokio::spawn(async move {
                let clear_fut = manager.clear_unused();
                let update_fut = manager.update_playlist();

                let (clear, update) = tokio::join!(clear_fut, update_fut);

                if let Err(e) = clear {
                    error!("Failed to clear unused files: {e}");
                }

                if let Err(e) = update {
                    error!("Failed to update playlist: {e}");
                }

            });
        }

        self.segment.write(data).await.map_err(|e| M3u8Error::SegmentError(e))?;

        Ok(())
    }

    pub async fn clear(self) -> Result<(), M3u8Error> {
        self.manager.clear().await?;

        Ok(())
    }
}
