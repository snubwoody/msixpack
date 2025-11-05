package bundle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse package", func(t *testing.T) {
		m := NewManifest()
		cfg := &Config{
			Package: ConfigPackage{
				Version:   "1.2.2.2",
				Publisher: "CN=Test",
			},
		}
		m.ParseConfig(cfg)
		i := m.Identity
		assert.Equal(t, i.Publisher, cfg.Package.Publisher)
		assert.Equal(t, i.Version, cfg.Package.Version)
	})

	t.Run("inherit id from name", func(t *testing.T) {
		m := NewManifest()
		cfg := &Config{
			Application: ConfigApplication{
				Name:        "Name",
				Description: "Description",
				Executable:  "file.exe",
			},
		}
		m.ParseConfig(cfg)
		a := m.Applications[0]
		assert.Equal(t, a.Id, cfg.Application.Name)
	})

	t.Run("parse application", func(t *testing.T) {
		m := NewManifest()
		cfg := &Config{
			Application: ConfigApplication{
				Id:          "Folio",
				Name:        "Name",
				Description: "Description",
				Executable:  "file.exe",
			},
		}
		m.ParseConfig(cfg)
		a := m.Applications[0]
		assert.Equal(t, a.Executable, cfg.Application.Executable)
		assert.Equal(t, a.Id, cfg.Application.Id)
		assert.Equal(t, a.EntryPoint, "Windows.FullTrustApplication")
	})
}

func TestDefaultPackageValues(t *testing.T) {
	m := NewManifest()
	assert.Equal(t, m.IgnorableNamespaces, "uap rescap")
	assert.Equal(t, m.ResCapNamespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities")
	assert.Equal(t, m.UapNamespace, "http://schemas.microsoft.com/appx/manifest/uap/windows10")
	assert.Equal(t, m.Namespace, "http://schemas.microsoft.com/appx/manifest/foundation/windows10")
}
