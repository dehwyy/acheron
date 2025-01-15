use tracing::Level;

pub struct LoggerConfig {
    pub log_level: Level,
    pub with_sentry: bool,
    pub with_target: bool,
    pub with_file: bool,
    pub with_line_number: bool,
    pub with_ansi: bool,
}

impl LoggerConfig {
    pub fn new() -> Self {
        Self::default()
    }

    pub fn new_production() -> Self {
        Self {
            log_level: Level::ERROR,
            with_sentry: true,
            with_ansi: false,
            with_file: true,
            with_target: true,
            with_line_number: true,
        }
    }

    pub fn sentry(mut self, enable: bool) -> Self {
        self.with_sentry = enable;
        self
    }
}

impl Default for LoggerConfig {
    fn default() -> Self {
        LoggerConfig {
            log_level: Level::INFO,
            with_sentry: true,
            with_ansi: true,
            with_file: true,
            with_target: true,
            with_line_number: true,
        }
    }
}
