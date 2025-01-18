use std::fmt::Display;


#[derive(Debug)]
pub enum SegmentFileError {
    CreateError,
    WriteError,
    FlushError,
}

impl Display for SegmentFileError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            SegmentFileError::CreateError => write!(f, "Failed to create segment file!"),
            SegmentFileError::WriteError => write!(f, "Failed to write segment file!"),
            SegmentFileError::FlushError => write!(f, "Failed to flush segment file!"),
        }
    }
}
