mod server;

use server::Server;

mod m3u8;

use log::{Logger, error, info};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    Logger::with_config(log::LoggerConfig::new().sentry(false));

    let addr_cfg = config::new::<config::Addr>();
    let port = addr_cfg.ports().srt_server;

    let m3u8_config = config::new::<config::M3u8>();

    let srv = Server::new(m3u8_config);

    match tokio::try_join!(srv.start(port),) {
        Ok((_,)) => info!("All done!"),
        Err(e) => error!("Error: {}", e),
    }

    Ok(())
}
