package bundle

import (
    "encoding/xml"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "os/exec"
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

// BundleApp bundles an folder with an appxmanifest.xml file into
// an msix package
func BundleApp(path string, output string) error {
    fmt.Println("Bundling app")
    cmd := exec.Command("./.msixpack/windows-toolkit/makeappx.exe", "pack", "/d", path, "/p", output, "/o")
    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("Output: %s\n", out)
        return err
    }
    fmt.Printf("Successfully created package -> %s\n", output)
    return nil
}

func createPackage() {
    pkg := Manifest{
        Namespace:       "http://schemas.microsoft.com/appx/manifest/foundation/windows10",
        UapNamespace:    "http://schemas.microsoft.com/appx/manifest/uap/windows10",
        ResCapNamespace: "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities",
        Properties: Properties{
            PublisherDisplayName: "Youtube",
        },
    }
    output, err := xml.MarshalIndent(pkg, "", "\t")
    if err != nil {
        fmt.Printf("Error: %s", err)
    }
    fmt.Printf("%s", output)
    err = os.WriteFile("appxmanifest.xml", output, 0666)
    if err != nil {
        //return err
    }
}
