mod srt_server;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr_cfg = config::new::<config::Addr>();
    let port = addr_cfg.ports().srt_server;

    match tokio::try_join!(srt_server::start(port),) {
        Ok((_,)) => println!("All done!"),
        Err(e) => eprintln!("Error: {}", e),
    }

    Ok(())
}
