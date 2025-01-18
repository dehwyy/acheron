use std::str::from_utf8;

use serde::de::DeserializeOwned;

#[allow(private_bounds)]
pub mod deserializable;
#[allow(private_bounds)]
pub mod env;

pub mod path;
use path::Path;

pub(crate) enum ConfigFormat {
    Env,
    TOML,
}

pub(crate) trait Parsable {
    const FILEPATH: Path<'_> = Path::new("config", "config.toml");
    const FORMAT: ConfigFormat = ConfigFormat::TOML;

    fn from(data: Vec<u8>) -> Self
    where
        Self: DeserializeOwned,
    {
        match Self::FORMAT {
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
        from_utf8(&data).expect(&format!("Failed to parse {} as .toml to &str.", T::FILEPATH)),
    )
    .expect("Failed to parse to toml.")
}
