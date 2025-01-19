pub use tracing::{debug, error, info, trace, warn};
use tracing_subscriber::FmtSubscriber;
use tracing_subscriber::fmt::time::ChronoLocal;
use tracing_subscriber::prelude::*;

mod config;
pub use config::LoggerConfig;

pub struct Logger;

impl Logger {
    pub fn initialize_with_config(config: LoggerConfig) {
        let fmt_subscriber = FmtSubscriber::builder()
            .pretty()
            .with_target(config.with_target)
            .with_file(config.with_file)
            .with_line_number(config.with_line_number)
            .with_ansi(config.with_ansi)
            .with_timer(ChronoLocal::new("%H:%M:%S".to_owned()))
            .with_max_level(config.log_level)
            .finish();

        match config.with_sentry {
            true => fmt_subscriber.with(sentry::integrations::tracing::layer()).init(),
            false => fmt_subscriber.init(),
        }
    }

    pub fn initialize() {
        let is_prod = std::env::var("PRODUCTION").map(|v| v == "TRUE").unwrap_or(false);

        let cfg = match is_prod {
            true => LoggerConfig::new_production(),
            false => LoggerConfig::new(),
        };

        Logger::initialize_with_config(cfg)
    }
}
