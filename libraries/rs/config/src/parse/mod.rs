use std::fs::File;
use std::io::Read;
use std::path::PathBuf;
use std::str::from_utf8;

use serde::de::DeserializeOwned;

pub(crate) trait Parsable {
    const filename: &'static str;

    fn from(data: Vec<u8>) -> Self
    where
        Self: DeserializeOwned,
    {
        toml::from_str(
            from_utf8(&data).expect(&format!("Failed to parse {} to &str.", Self::filename)),
        )
        .expect("Failed to parse to toml.")
    }
}

pub fn new<T>() -> T
where
    T: Parsable + DeserializeOwned,
{
    T::from(read_file::<T>())
}

fn read_file<T: Parsable>() -> Vec<u8> {
    let mut path = PathBuf::from("config");
    path.push(T::filename);

    let mut filebuf = vec![];
    File::open(path)
        .expect("Failed to open config file.")
        .read_to_end(&mut filebuf)
        .expect("Failed to read config file.");

    filebuf
}
