mod server;
use std::env;

use server::Server;

mod m3u8;

use log::{Logger, error, info};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    Logger::with_config(log::LoggerConfig::new().sentry(false));

    let addr_cfg = config::new::<config::Addr>();
    let port = addr_cfg.ports().srt_server;

    let mut args = env::args();
    let srv = Server::new(
        args.nth(1).map(|v| v.parse::<u64>().unwrap()).unwrap_or(2000),
        args.nth(0).map(|v| v.parse::<usize>().unwrap()).unwrap_or(15),
    );

    match tokio::try_join!(srv.start(port),) {
        Ok((_,)) => info!("All done!"),
        Err(e) => error!("Error: {}", e),
    }

    Ok(())
}
