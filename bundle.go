package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Manifest struct {
	XMLName         xml.Name `xml:"Package"`
	Namespace       string   `xml:"xmlns,attr"`
	UapNamespace    string   `xml:"xmlns:uap,attr"`
	ResCapNameSpace string   `xml:"xmlns:rescap,attr"`
	Properties      Properties
}

type Properties struct {
	DisplayName          string `xml:"DisplayName"`
	PublisherDisplayName string `xml:"PublisherDisplayName"`
	Logo                 string `xml:"Logo"`
}

type Identity struct {
}

// Create a new [Manifest].
func NewManifest() *Manifest {
	// These values should never change.
	m := &Manifest{
		Namespace:       "http://schemas.microsoft.com/appx/manifest/foundation/windows10",
		UapNamespace:    "http://schemas.microsoft.com/appx/manifest/uap/windows10",
		ResCapNameSpace: "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities",
	}

	return m
}

func createPackage() {
	pkg := Manifest{
		Namespace:       "http://schemas.microsoft.com/appx/manifest/foundation/windows10",
		UapNamespace:    "http://schemas.microsoft.com/appx/manifest/uap/windows10",
		ResCapNameSpace: "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities",
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
