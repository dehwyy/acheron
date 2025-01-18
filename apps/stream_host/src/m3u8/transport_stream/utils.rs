use std::path::PathBuf;

use tokio::fs::{self, File};

pub fn get_filename(segment_index: usize) -> String {
    format!("segment_{}.ts", segment_index)
}

pub async fn remove_segment(
    directory: PathBuf,
    segment_index: usize,
) -> Result<(), std::io::Error> {
    fs::remove_file(directory.join(get_filename(segment_index))).await
}

pub(super) async fn create_segment_file(
    directory: PathBuf,
    idx: usize,
) -> Result<File, std::io::Error> {
    File::create(directory.join(get_filename(idx))).await
}
