package bundle

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/exec"
	"path"
)

type Manifest struct {
	XMLName            xml.Name `xml:"Package"`
	Namespace          string   `xml:"xmlns,attr"`
	UapNamespace       string   `xml:"xmlns:uap,attr"`
	ResCapNamespace    string   `xml:"xmlns:rescap,attr"`
	IgnorableNamespace string   `xml:"IgnorableNamespace,attr"`
	Properties         Properties
	Identity           Identity
	Dependencies       Dependencies
	Applications       []Application `xml:"Applications>Application"`
	Capabilities       []Capability  `xml:"Capabilities>rescap:Capability"`
}

type Capability struct {
	Name string `xml:"Name,attr"`
}

type Application struct {
	Id         string `xml:"Id,attr"`
	Executable string `xml:"Executable,attr"`
	EntryPoint string `xml:"EntryPoint,attr"`
}

type Properties struct {
	DisplayName          string `xml:"DisplayName"`
	PublisherDisplayName string `xml:"PublisherDisplayName"`
	Logo                 string `xml:"Logo"`
}

type Identity struct {
	Name                  string `xml:"Name,attr"`
	Version               string `xml:"Version,attr"`
	Publisher             string `xml:"Publisher,attr"`
	ProcessorArchitecture string `xml:"ProcessorArchitecture,attr"`
}

type Dependencies struct {
	TargetDeviceFamily TargetDeviceFamily
}

type TargetDeviceFamily struct {
	Name             string `xml:"Name,attr"`
	MinVersion       string `xml:"MinVersion,attr"`
	MaxVersionTested string `xml:"MaxVersionTested,attr"`
}

// Create a new [Manifest].
func NewManifest() *Manifest {
	// These values should never change.
	m := &Manifest{
		Namespace:          "http://schemas.microsoft.com/appx/manifest/foundation/windows10",
		UapNamespace:       "http://schemas.microsoft.com/appx/manifest/uap/windows10",
		ResCapNamespace:    "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities",
		IgnorableNamespace: "uap rescap",
	}

	return m
}

func LoadConfig(m *Manifest) error {
	// TODO: maybe display-name
	displayName := viper.GetString("name")
	publisherName := viper.GetString("publisher-name")
	logo := viper.GetString("logo")
	m.Properties = Properties{
		DisplayName:          displayName,
		PublisherDisplayName: publisherName,
		Logo:                 logo,
	}
	loadIdentity(m)
	loadDependencies(m)
	loadApplication(m)
	loadCapabilties(m)
	return nil
}

func loadIdentity(m *Manifest) {
	name := viper.GetString("identity-name")
	version := viper.GetString("version")
	publisher := viper.GetString("publisher")
	arch := viper.GetString("architecture")
	m.Identity = Identity{
		Name:                  name,
		Version:               version,
		Publisher:             publisher,
		ProcessorArchitecture: arch,
	}
}

func loadDependencies(m *Manifest) {
	t := TargetDeviceFamily{
		Name:             "Windows.Desktop",
		MinVersion:       "10.0.17763.0",
		MaxVersionTested: "10.0.22621.0",
	}

	m.Dependencies = Dependencies{
		TargetDeviceFamily: t,
	}
}

func loadApplication(m *Manifest) {
	m.Applications = []Application{
		{
			Id:         "Youtube",
			Executable: "youtube.exe",
			EntryPoint: "Windows.FullTrustApplication",
		},
	}
}

func loadCapabilties(m *Manifest) {
	m.Capabilities = []Capability{
		{
			Name: "runFullTrust",
		},
	}
}

func ValidateToolkit() error {
	// TODO: check executions
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dest := path.Join(home, ".msixpack")
	err = os.MkdirAll(dest, 0666)
	if err != nil {
		return err
	}
	p, err := downloadToolkit(dest)
	if err != nil {
		return err
	}

	err = UnzipFile(p, dest)
	if err != nil {
		return err
	}
	// TODO: get the name of the destination folder
	err = CopyDir(path.Join(dest, "MSIX-Toolkit-2.0/WindowsSDK/11/10.0.22000.0/x64"), path.Join(dest, "windows-toolkit"))
	if err != nil {
		return err
	}
	return nil
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
		fmt.Printf("%s\n", file.Name())
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()
		io.Copy(destFile, srcFile)
	}
	return nil
}

func UnzipFile(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

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
	}
	return nil
}

func CopyFile() {

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
	cmd := exec.Command(bin, "pack", "/d", p, "/p", output, "/o")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Output: %s\n", out)
		return err
	}
	fmt.Printf("Successfully created package -> %s\n", output)
	return nil
}
