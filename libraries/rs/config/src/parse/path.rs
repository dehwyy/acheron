#[derive(Debug)]
pub(crate) struct Path<'a> {
    dir: Option<&'a str>,
    filename: &'a str,
}

impl<'a> Path<'_>
where
    'a: 'static,
{
    pub const fn new_file(filename: &'a str) -> Self {
        Self {
            dir: None,
            filename,
        }
    }

    pub const fn new(dir: &'a str, filename: &'a str) -> Self {
        Self {
            dir: Some(dir),
            filename,
        }
    }
}

impl std::fmt::Display for Path<'_> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut path = std::path::PathBuf::new();
        if let Some(dir) = self.dir {
            path.push(dir);
        }
        path.push(self.filename);

        write!(f, "{}", path.to_str().expect("Invalid path. (Cannot decode to UTF-8)."))
    }
}
