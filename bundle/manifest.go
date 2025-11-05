package bundle

import (
	"encoding/xml"
)

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
	// The namespace values should never change.
	m := &Manifest{
		Namespace:          "http://schemas.microsoft.com/appx/manifest/foundation/windows10",
		UapNamespace:       "http://schemas.microsoft.com/appx/manifest/uap/windows10",
		ResCapNamespace:    "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities",
		IgnorableNamespace: "uap rescap",
		Dependencies: Dependencies{
			TargetDeviceFamily: TargetDeviceFamily{
				Name:             "Windows.Desktop",
				MinVersion:       "10.0.17763.0",
				MaxVersionTested: "10.0.22621.0",
			},
		},
		Capabilities: []Capability{
			{
				Name: "runFullTrust",
			},
		},
	}
	return m
}

func (m *Manifest) ParseConfig(cfg *Config) {
	m.Properties = Properties{
		DisplayName:          cfg.Package.DisplayName,
		PublisherDisplayName: cfg.Package.PublisherName,
		Logo:                 cfg.Package.Logo,
	}

	m.Identity = Identity{
		Name:                  cfg.Package.Name,
		Version:               cfg.Package.Version,
		Publisher:             cfg.Package.Publisher,
		ProcessorArchitecture: "x64",
	}

	m.loadApplication(cfg)
}

func (m *Manifest) loadApplication(cfg *Config) {
	id := cfg.Application.Id
	if cfg.Application.Id == "" {
		id = cfg.Application.Name
	}
	m.Applications = []Application{
		{
			Id:         id,
			Executable: cfg.Application.Executable,
			EntryPoint: "Windows.FullTrustApplication",
		},
	}
}
