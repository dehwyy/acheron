use std::fmt::Display;
use std::io;

use super::transport_stream::SegmentFileError;


pub enum M3u8Error {
    InitializationError(io::Error),
    ClearError(io::Error),

    OpenPlaylistFileError(io::Error),
    WritePlaylistFileError,

    SegmentError(SegmentFileError),
}

impl Display for M3u8Error {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            M3u8Error::InitializationError(e) => write!(f, "Initialization error: {e}!"),
            M3u8Error::ClearError(e) => write!(f, "Clear error: {e}!"),
            M3u8Error::OpenPlaylistFileError(e) => write!(f, "Open playlist file error: {e}!"),
            M3u8Error::WritePlaylistFileError => write!(f, "Write playlist file error!"),
            M3u8Error::SegmentError(e) => write!(f, "Segment error: {e}!"),
        }
    }
}
