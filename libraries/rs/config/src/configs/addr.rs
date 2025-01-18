use serde::Deserialize;

#[derive(Deserialize)]
pub struct Ports {
    pub srt_server: u16,
}

#[derive(Deserialize)]
pub struct Addr {
    ports: Ports,
}

impl Addr {
    pub fn ports(&self) -> &Ports {
        &self.ports
    }
}

impl crate::parse::Parsable for Addr {}
