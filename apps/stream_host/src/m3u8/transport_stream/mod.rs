use std::path::PathBuf;
use std::time::Instant;

use tokio::fs::File;
use tokio::io::AsyncWriteExt;

pub mod utils;

mod errors;
pub use errors::*;

#[derive(Clone)]
pub struct SegmentMetadata {
    /// Where to store files
    pub directory: PathBuf,

    /// Current write segment index
    pub segment_index: usize,

    /// Last write timestamp
    pub last_write: Instant,
}

impl SegmentMetadata {
    pub fn new(directory: PathBuf) -> Self {
        Self {
            directory,
            segment_index: 0,
            last_write: Instant::now(),
        }
    }
}

pub struct SegmentFile {
    directory: PathBuf,
    file: File,
}

impl SegmentFile {
    pub async fn try_new(directory: PathBuf) -> Result<Self, SegmentFileError> {
        Ok(Self {
            file: utils::create_segment_file(directory.clone(), 0)
                .await
                .map_err(|_| SegmentFileError::CreateError)?,
            directory,
        })
    }

    pub async fn try_next(&mut self, idx: &mut usize) -> Result<(), SegmentFileError> {
        self.file.flush().await.map_err(|_| SegmentFileError::FlushError)?;

        *idx += 1;

        self.file = utils::create_segment_file(self.directory.clone(), *idx)
            .await
            .map_err(|_| SegmentFileError::CreateError)?;

        Ok(())
    }

    pub async fn write(&mut self, data: &[u8]) -> Result<(), SegmentFileError> {
        self.file.write_all(data).await.map_err(|_| SegmentFileError::WriteError)
    }
}
