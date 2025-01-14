use std::ops::Deref;

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

// Polymorphism in Rust hahaha
impl Deref for Addr {
    type Target = Ports;
    fn deref(&self) -> &Self::Target {
        &self.ports
    }
}

impl crate::parse::Parsable for Addr {}
