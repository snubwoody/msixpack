package bundle

import (
	"encoding/xml"
	"github.com/spf13/viper"
)

type ManifestConfig struct {
	// The application id, this value only needs to be unique
	// inside the package but not globally, this value
	// must not be changed after the app has been published to the
	// microsoft store.
	//
	// If this value is empty then it will be inherited from
	// the name/
	Id string
	// The name of the application
	Name string
	// The description of the app.
	Description string
	// The path to the executable.
	Executable string
	// Describes the publisher information.
	// The Publisher attribute must match the publisher subject information of
	// the certificate used to sign a package.
	Publisher string
	// The version of the package, must be in quad notation,
	// MAJOR.MINOR.PATCH.BUILD
	//
	// The major version cannot be 0.
	Version string
}

// TODO: maybe viper isnt needed
// TODO: add visual elements
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
	c := &ManifestConfig{}
	loadIdentity(m, c)
	loadDependencies(m)
	loadApplication(m, c)
	loadCapabilties(m)
	return nil
}

func loadIdentity(m *Manifest, cfg *ManifestConfig) {
	name := viper.GetString("identity-name")

	m.Identity = Identity{
		Name:                  name,
		Version:               cfg.Version,
		Publisher:             cfg.Publisher,
		ProcessorArchitecture: "x64",
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

func loadApplication(m *Manifest, cfg *ManifestConfig) {
	id := cfg.Id
	if cfg.Id == "" {
		id = cfg.Name
	}
	m.Applications = []Application{
		{
			Id:         id,
			Executable: cfg.Executable,
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
