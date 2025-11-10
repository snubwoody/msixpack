use std::fs;
use std::fs::File;
use std::path::{Path, PathBuf};
use anyhow::Context;
use glob::glob;
use serde::{Deserialize, Serialize};

fn main() -> anyhow::Result<()> {
    // TODO:
    // Copy executable and resources
    // Create appxmanifest
    // Create msix package
    let config_path = Path::new("testdata/hello/msixpack.toml");
    let bytes = fs::read(config_path)?;
    let mut config: Config = toml::from_slice(&bytes)?;
    config.directory = config_path.parent().unwrap().to_path_buf();
    dbg!(&config);
    create_package(&config,"out-temp")?;
    Ok(())
}

#[derive(Debug,Clone,PartialEq,Serialize,Deserialize,Default)]
struct Config{
    /// The path of the configuration file.
    #[serde(skip,default)]
    directory: PathBuf,
    package: Package,
    application: Application
}

#[derive(Debug,Clone,PartialEq,Serialize,Deserialize,Default)]
struct Package{
    /// A series of glob paths.
    resources: Vec<String>
}

#[derive(Debug,Clone,PartialEq,Serialize,Deserialize,Default)]
struct Application{
    executable: PathBuf,
}

fn create_package(config: &Config,dest: impl AsRef<Path>) -> anyhow::Result<()> {
    // FIXME: create dest directory

    let dest = dest.as_ref();
    let exe_path = config.directory.join(&config.application.executable);
    let exe = config.application.executable.file_name().unwrap();
    // FIXME: put it in the root
    fs::copy(exe_path, dest.join(&exe))?;
    copy_resources(config, &dest)?;
    Ok(())
}

/// Copy all the resources defined in the [`Config`] to the destination directory.
fn copy_resources(config: &Config,dest: impl AsRef<Path>) -> anyhow::Result<()> {
    let dir = &config.directory;
    for pattern in &config.package.resources{
        let path = dir.join(pattern);
        for entry in glob(path.to_str().unwrap())?{
            let entry = entry?;
            let base_path = entry.strip_prefix(&dir)?;
            let output = dest.as_ref().join(base_path);
            // FIXME: this will panic for base patterns
            fs::create_dir_all(&output.parent().unwrap())?;
            fs::copy(entry, output)?;
        }
    }
    Ok(())
}

/// Copy all glob matches into the destination directory.
fn copy_glob(){

}

#[cfg(test)]
mod test{
    use std::fs::File;
    use tempfile::{tempdir, tempfile};
    use super::*;

    #[test]
    fn copy_resources() -> anyhow::Result<()> {
        let dir = tempdir()?;
        let icons_dir = dir.path().join("icons");
        fs::create_dir(&icons_dir)?;
        let icon1 = dir.path().join("icons").join("icon1.png");
        let icon2 = dir.path().join("icons").join("icon2.png");
        File::create(&icon1)?;
        File::create(&icon2)?;
        let config = Config{
            directory: dir.path().to_path_buf(),
            package: Package{
                resources: vec!["icons/*.png".to_string()]
            },
            ..Default::default()
        };
        let out = tempdir()?;
        super::copy_resources(&config,out.path())?;
        assert!(fs::exists(out.path().join("icons/icon1.png"))?);
        assert!(fs::exists(out.path().join("icons/icon2.png"))?);
        Ok(())
    }
}
