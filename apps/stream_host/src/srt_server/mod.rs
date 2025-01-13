use futures::{StreamExt, TryStreamExt};
use srt_protocol::settings::KeySettings;
use srt_tokio::{ConnectionRequest, SrtListener, SrtSocket};
use tokio::net::{TcpListener, TcpStream};

pub async fn start(port: u16) -> Result<(), Box<dyn std::error::Error>> {
    let (_server, mut incoming) = SrtListener::builder().bind(port).await?;

    while let Some(connection_req) = incoming.incoming().next().await {
        tokio::spawn(handle_srt_connection(connection_req));
    }

    Ok(())
}

async fn handle_srt_connection(conn: ConnectionRequest) {
    // Todo: perform some validation / authentication.
    let key_settings: Option<KeySettings> = None;

    let remote_addr = conn.remote();
    println!("New connection: {remote_addr}");

    let mut socket = conn.accept(key_settings).await.expect("Error accepting connection.");

    let mut buf: Vec<u8> = vec![];
    let mut i = 0;
    while let Ok(Some((timestamp, data))) = socket.try_next().await {
        println!(
            "{remote_addr}: frame={i} timestamp={timestamp:?} datasize={} bufsize(kb)={}",
            data.len(),
            buf.len() / 1024
        );
        i += 1;
        buf.extend_from_slice(&data);
    }
}
