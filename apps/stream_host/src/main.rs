use log::{error, info};
mod srt_server;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    log::Logger::with_config(log::LoggerConfig::new().sentry(false));

    let addr_cfg = config::new::<config::Addr>();
    let port = addr_cfg.ports().srt_server;

    match tokio::try_join!(srt_server::start(port),) {
        Ok((_,)) => info!("All done!"),
        Err(e) => error!("Error: {}", e),
    }

    Ok(())
}
