use std::path::PathBuf;
use std::sync::Arc;

use futures::{StreamExt, TryStreamExt};
use log::{error, info};
use srt_protocol::settings::KeySettings;
use srt_tokio::{ConnectionRequest, SrtListener, SrtSocket};


use crate::m3u8::{M3u8, M3u8Config};

// ? Maybe useful: "#EXT-X-DISCONTINUITY" @gpt "Use if switching codecs, resolutions, or timestamp

mod errors;
use errors::*;

pub enum ConnectionStatus {
    ConnectionClosed,
}

pub struct Server {
    m3u8_config: M3u8Config,
}

impl Server {
    pub fn new(m3u8_config: config::M3u8) -> Self {
        Self {
            m3u8_config: M3u8Config {
                segment_duration: std::time::Duration::from_millis(
                    m3u8_config.segment_duration_millis,
                ),
                segments_per_stream: m3u8_config.segments_per_stream,
                clear_unused_segments: m3u8_config.clear_after_unused_segments,
                streams_directory: PathBuf::from(m3u8_config.streams_directory.clone()),
            },
        }
    }

    pub async fn start(self, port: u16) -> Result<(), Box<dyn std::error::Error>> {
        let (_server, mut incoming) = SrtListener::builder().bind(port).await?;

        info!("Listening on port {}", port);

        let shared = Arc::new(self);
        while let Some(connection_req) = incoming.incoming().next().await {
            let shared_clone = shared.clone();
            tokio::spawn(async move {
                match shared_clone.handle_connection(connection_req).await {
                    Ok(status) => match status {
                        ConnectionStatus::ConnectionClosed => info!("Connection closed"),
                    },
                    Err(err) => {
                        error!("Connection error: {}", err);
                    }
                }
            });
        }

        Ok(())
    }

    async fn authorize_connecton(
        conn: ConnectionRequest,
    ) -> Result<(SrtSocket, String), AuthorizationError> {
        // TODO: perform auth.
        let key_settings: Option<KeySettings> = None;

        let remote_addr = conn.remote();
        let stream_id =
            conn.stream_id().ok_or(AuthorizationError::StreamIdNotProvided)?.to_string();

        info!("New connection: {remote_addr}");
        let socket = conn
            .accept(key_settings)
            .await
            .map_err(|_| AuthorizationError::FailedToAcceptConnection)?;

        Ok((socket, stream_id))
    }

    async fn handle_connection(
        self: Arc<Self>,
        conn: ConnectionRequest,
    ) -> Result<ConnectionStatus, ConnectionError> {
        let (mut socket, stream_id) = Self::authorize_connecton(conn)
            .await
            .map_err(|e| ConnectionError::AuthorizationError(e))?;

        let mut playlist = M3u8::try_new(self.m3u8_config.clone(), stream_id)
            .await
            .map_err(|e| ConnectionError::M3u8Error(e))?;

        loop {
            match socket.try_next().await {
                Ok(Some((timestamp, data))) => {
                    if let Err(err) = playlist.handle(timestamp, &data).await {
                        error!("Failed to handle data: {err}");
                    };
                },
                Err(_err) => {
                    return Err(ConnectionError::InternalError);
                },
                Ok(None) => {
                    playlist.clear().await.map_err(|e| ConnectionError::M3u8Error(e))?;
                    return Ok(ConnectionStatus::ConnectionClosed);
                },
            }
        }
    }
}
