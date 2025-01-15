## Logger
- Call either `new` or `with_config` to initialize logger.
- Import `info!`, `error!`, ***etc.*** to log something (convenient macros)

## LoggerConfig
- ***log_level***: Sets the maximum [verbosity level] that will be enabled by the subscriber. Levels from higher(less verbose) to lower(most verbose): ***Error->Warn->Info->Debug->Trace***
- ***sentry***: Enables `sentry` integration
- ***ansi***: If the formatter emits ANSI terminal escape codes for colors and other text formatting.
- ***file***: If event's source code file path is displayed.
- ***target***: If event's target is displayed.
- ***line_number***: If event's source code line number is displayed.
