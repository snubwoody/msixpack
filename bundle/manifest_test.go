package bundle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadIdentity(t *testing.T) {
	t.Run("load config values", func(t *testing.T) {
		m := NewManifest()
		cfg := &ManifestConfig{
			Version:   "1.2.2.2",
			Publisher: "CN=Test",
		}
		loadApplication(m, cfg)
		i := m.Identity
		assert.Equal(t, i.Publisher, cfg.Publisher)
		assert.Equal(t, i.Version, cfg.Version)
	})

	t.Run("inherit id from name", func(t *testing.T) {
		m := NewManifest()
		cfg := &ManifestConfig{
			Name:        "Name",
			Description: "Description",
			Executable:  "file.exe",
		}
		loadApplication(m, cfg)
		a := m.Applications[0]
		assert.Equal(t, a.Id, cfg.Name)
	})
}

func TestLoadApplication(t *testing.T) {
	t.Run("load config values", func(t *testing.T) {
		m := NewManifest()
		cfg := &ManifestConfig{
			Id:          "Folio",
			Name:        "Name",
			Description: "Description",
			Executable:  "file.exe",
		}
		loadApplication(m, cfg)
		a := m.Applications[0]
		assert.Equal(t, a.Executable, cfg.Executable)
		assert.Equal(t, a.Id, cfg.Id)
		assert.Equal(t, a.EntryPoint, "Windows.FullTrustApplication")
	})

	t.Run("inherit id from name", func(t *testing.T) {
		m := NewManifest()
		cfg := &ManifestConfig{
			Name:        "Name",
			Description: "Description",
			Executable:  "file.exe",
		}
		loadApplication(m, cfg)
		a := m.Applications[0]
		assert.Equal(t, a.Id, cfg.Name)
	})
}

func TestDefaultPackageValues(t *testing.T) {
	m := NewManifest()
	assert.Equal(t, m.IgnorableNamespace, "uap rescap")
	assert.Equal(t, m.ResCapNamespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities")
	assert.Equal(t, m.UapNamespace, "http://schemas.microsoft.com/appx/manifest/uap/windows10")
	assert.Equal(t, m.Namespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10")
}
