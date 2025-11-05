package bundle

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

func ValidateToolkit() error {
	// TODO: check executions
	// TODO: check if file already exists
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	_, err = os.Stat(path.Join(home, ".msixpack", "windows-toolkit", "makeappx.exe"))
	if !os.IsNotExist(err) {
		return nil
	}

	dest := path.Join(home, ".msixpack")
	err = os.MkdirAll(dest, 0666)
	if err != nil {
		return err
	}
	zipFile, err := downloadToolkit(dest)
	if err != nil {
		return err
	}

	err = UnzipFile(zipFile, dest)
	if err != nil {
		return err
	}
	// TODO: get the name of the destination folder
	toolkitPath := path.Join(dest, "windows-toolkit")
	err = CopyDir(path.Join(dest, "MSIX-Toolkit-2.0/WindowsSDK/11/10.0.22000.0/x64"), toolkitPath)
	if err != nil {
		return err
	}
	fmt.Printf("Installed windows toolkit at %s\n", toolkitPath)
	// Delete the zip file and unused contents
	err = os.RemoveAll(zipFile)
	if err != nil {
		return err
	}
	return os.RemoveAll(path.Join(dest, "MSIX-Toolkit-2.0"))
}

// CopyDir copies all the files in src into the directory dest.
func CopyDir(src string, dest string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := path.Join(src, file.Name())
		destPath := path.Join(dest, file.Name())
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		if _, err = io.Copy(destFile, srcFile); err != nil {
			return err
		}
		if err = srcFile.Close(); err != nil {
			return err
		}
		if err = destFile.Close(); err != nil {
			return err
		}
	}
	return nil
}

func UnzipFile(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		out := path.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(out, 0777)
			if err != nil {
				return err
			}
			continue
		}

		//fmt.Printf("Unzipping %s\n", out)
		newFile, err := os.Create(out)
		if err != nil {
			return err
		}
		_, err = io.Copy(newFile, rc)
		if err != nil {
			return err
		}
		if err := newFile.Close(); err != nil {
			return err
		}
		if err := rc.Close(); err != nil {
			return err
		}
	}
	return r.Close()
}

// BundleApp bundles an folder with an appxmanifest.xml file into
// an msix package
func BundleApp(p string, output string) error {
	fmt.Println("Bundling app")
	err := ValidateToolkit()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	bin := path.Join(home, ".msixpack", "windows-toolkit", "makeappx.exe")
	err = packMsix(bin, p, output)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully created package -> %s\n", output)
	return nil
}

func packMsix(makeappxPath, dir, output string) error {
	if stat, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("the system cannot find the directory: %s", dir)
	} else if err != nil {
		return err
	} else if !stat.IsDir() {
		return fmt.Errorf("the input package must be a directory")
	}

	cmd := exec.Command(makeappxPath, "pack", "/d", dir, "/p", output, "/o")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Output: %s\n", out)
		return err
	}
	return err
}
