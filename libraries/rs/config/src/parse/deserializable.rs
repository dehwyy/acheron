use std::fs::File;
use std::io::Read;

use serde::de::DeserializeOwned;

use super::Parsable;

trait Deserializable: Parsable + DeserializeOwned {}
impl<T: Parsable + DeserializeOwned> Deserializable for T {}

/// Reads the config file and returns the parsed config.
pub fn new<T>() -> T
where
    T: Deserializable,
{
    T::from(read_file::<T>())
}

fn read_file<T: Parsable>() -> Vec<u8> {
    let mut filebuf = vec![];
    File::open(T::FILEPATH.to_string())
        .expect("Failed to open config file.")
        .read_to_end(&mut filebuf)
        .expect("Failed to read config file.");

    filebuf
}
