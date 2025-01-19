use std::ops::Deref;

use serde::Deserialize;

#[derive(Deserialize)]
pub struct M3u8Config {
    pub segment_duration_millis: u64,
    pub segments_per_stream: usize,
    pub clear_after_unused_segments: Option<usize>,
    pub streams_directory: String,
}

#[derive(Deserialize)]
pub struct M3u8 {
    pub m3u8: M3u8Config,
}

// Polymorphism in Rust hahaha
impl Deref for M3u8 {
    type Target = M3u8Config;
    fn deref(&self) -> &Self::Target {
        &self.m3u8
    }
}

impl crate::parse::Parsable for M3u8 {}
