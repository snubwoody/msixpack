use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Default)]
#[serde(rename="Package")]
pub struct AppxManifest{
    #[serde(rename = "@xmlns")]
    pub xmlns: String,
    #[serde(rename = "@xmlns:uap")]
    pub xmlns_uap: String,
    #[serde(rename = "@xmlns:rescap")]
    pub xmlns_rescap: String,
    #[serde(rename = "@IgnorableNamespaces")]
    pub ignorable_namespaces: String,
    pub identity: Identity
}

impl AppxManifest{
    /// Create a new appxmanifest.
    pub fn new() -> AppxManifest{
        Self{
            xmlns: String::from("http://schemas.microsoft.com/appx/manifest/foundation/windows10"),
            xmlns_uap: String::from("http://schemas.microsoft.com/appx/manifest/uap/windows10"),
            xmlns_rescap: String::from("http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities"),
            ignorable_namespaces: String::from("uap rescap"),
            ..Default::default()
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Default)]
pub struct Identity{
    #[serde(rename = "@xmlns:rescap")]
    pub name: String,
    pub version: String,
    pub publisher: String,
    pub processor_architecture: String,
}

#[cfg(test)]
mod test{
    use super::*;

    #[test]
    fn manifest_namespace(){
        let manifest = AppxManifest::new();
        assert_eq!(manifest.xmlns,"http://schemas.microsoft.com/appx/manifest/foundation/windows10");
        assert_eq!(manifest.xmlns_uap,"http://schemas.microsoft.com/appx/manifest/uap/windows10");
        assert_eq!(manifest.xmlns_rescap,"http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities");
        assert_eq!(manifest.ignorable_namespaces,"uap rescap");
    }
}
