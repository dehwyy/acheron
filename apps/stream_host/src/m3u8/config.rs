use std::{path::PathBuf, time::Duration};

#[derive(Clone, Debug)]
pub struct M3u8Config {
    /// Duration of one segment. Tip: should be greater than S second, where S is keyframe
    /// interval.
    pub segment_duration: Duration,

    // How many segments should be stored in `playlist`.
    pub segments_per_stream: usize,

    /// Some(n) - `n` is difference between last unused and file to be deleted.
    /// None    - do not delete unused segments.
    pub clear_unused_segments: Option<usize>,

    /// "Where all streams folder should be located at."
    pub streams_directory: PathBuf,
}
