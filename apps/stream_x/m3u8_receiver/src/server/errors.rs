use std::fmt::Display;

use crate::m3u8::M3u8Error;

pub enum ConnectionError {
    AuthorizationError(AuthorizationError),
    M3u8Error(M3u8Error),
    InternalError,
}

pub enum AuthorizationError {
    StreamIdNotProvided,
    FailedToAcceptConnection,
}

impl Display for ConnectionError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            ConnectionError::AuthorizationError(e) => write!(f, "{e}"),
            ConnectionError::M3u8Error(e) => write!(f, "{e}"),
            ConnectionError::InternalError => write!(f, "Internal error!"),
        }
    }
}

impl Display for AuthorizationError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            AuthorizationError::StreamIdNotProvided => write!(f, "Stream ID not provided!"),
            AuthorizationError::FailedToAcceptConnection => {
                write!(f, "Failed to accept connection!")
            },
        }
    }
}
