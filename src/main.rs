use std::fs;
use std::fs::File;
use std::path::{Path, PathBuf};
use anyhow::Context;
use glob::glob;

fn main() {
    // TODO:
    // Copy executable and resources
    // Create appxmanifest
    // Create msix package
    println!("Hello, world!");
}

#[derive(Debug,Clone,PartialEq)]
struct Config{
    /// The path of the configuration file.
    directory: PathBuf,
    /// A series of glob paths.
    resources: Vec<String>
}

struct Application{

}

fn create_package(path: impl AsRef<Path>,out: impl AsRef<Path>) -> anyhow::Result<()> {
    fs::copy(path.as_ref(), out)
        .with_context(|| format!("Failed to copy file to {}", path.as_ref().display()))?;
    Ok(())
}

/// Copy all the resources defined in the [`Config`] to the destination directory.
fn copy_resources(config: &Config,dest: impl AsRef<Path>) -> anyhow::Result<()> {
    let dir = &config.directory;
    for pattern in &config.resources{
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
    fn copy_executable() -> anyhow::Result<()> {
        let dir = tempdir()?;
        let exe_path = dir.path().join("main.exe");
        File::create(&exe_path)?;
        let out_path = dir.path().join("out.exe");
        create_package(exe_path,&out_path)?;
        Ok(())
    }

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
            resources: vec!["icons/*.png".to_string()]
        };
        let out = tempdir()?;
        super::copy_resources(&config,out.path())?;
        assert!(fs::exists(out.path().join("icons/icon1.png"))?);
        assert!(fs::exists(out.path().join("icons/icon2.png"))?);
        Ok(())
    }
}
