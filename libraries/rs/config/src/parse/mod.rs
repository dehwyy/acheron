use std::str::from_utf8;

use serde::de::DeserializeOwned;

pub mod deserializable;
pub mod env;

pub mod path;
use path::Path;

pub(crate) enum ConfigFormat {
    Env,
    TOML,
}

pub(crate) trait Parsable {
    const filepath: Path<'_> = Path::new("config", "config.toml");
    const format: ConfigFormat = ConfigFormat::TOML;

    fn from(data: Vec<u8>) -> Self
    where
        Self: DeserializeOwned,
    {
        match Self::format {
            ConfigFormat::TOML => parse_toml(data),
            _ => {
                panic!(
                    "Unsupported config format. Maybe you should use `new_env` (e.g for `env` config) instead of `new` function."
                )
            },
        }
    }
}

fn parse_toml<T: Parsable + DeserializeOwned>(data: Vec<u8>) -> T {
    toml::from_str(
        from_utf8(&data).expect(&format!("Failed to parse {} as .toml to &str.", T::filepath)),
    )
    .expect("Failed to parse to toml.")
}
