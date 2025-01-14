use std::env;

use envconfig::Envconfig;

use crate::parse::ConfigFormat;
use crate::parse::path::Path;

#[derive(Envconfig)]
/// Config, which relies on environment variables.
pub struct Config {
    #[envconfig(from = "SENTRY_DSN", default = "")]
    pub sentry_dsn: String,
}

impl crate::parse::Parsable for Config {
    const filepath: Path<'_> = Path::new_file(".env");
    const format: ConfigFormat = ConfigFormat::Env;
}
