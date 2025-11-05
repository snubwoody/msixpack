package bundle

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadIdentity(t *testing.T) {
	viper.Set("name")
	m := NewManifest()
	loadIdentity(m)
	fmt.Printf("Manifest: %v\n", m)
}

func TestDefaultPackageValues(t *testing.T) {
	m := NewManifest()
	assert.Equal(t, m.IgnorableNamespace, "uap rescap")
	assert.Equal(t, m.ResCapNamespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities")
	assert.Equal(t, m.UapNamespace, "http://schemas.microsoft.com/appx/manifest/uap/windows10")
	assert.Equal(t, m.Namespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10")
}
