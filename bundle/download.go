package bundle

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const windowsToolkitUrl = "https://github.com/microsoft/MSIX-Toolkit/archive/refs/tags/v2.0.zip"

// Downloads the windows toolkit.
func downloadToolkit(basedir string) (string, error) {
	p := path.Join(basedir, "windows-toolkit.zip")
	fmt.Printf("Downloading windows toolkit from %s\n", windowsToolkitUrl)
	_, err := downloadFile(windowsToolkitUrl, p)
	if err != nil {
		return "", err
	}

	return p, nil
}

func downloadFile(url string, dest string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}
	if err = resp.Body.Close(); err != nil {
		return "", err
	}
	return dest, nil
}
