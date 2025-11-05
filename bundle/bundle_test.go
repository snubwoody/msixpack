package bundle

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestPackMsix(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)
	makeappxPath := path.Join(home, ".msixpack", "windows-toolkit", "makeappx.exe")
	tmp, err := os.MkdirTemp("", "test-msix")
	require.NoError(t, err)

	t.Run("directory does not exist", func(t *testing.T) {
		err = packMsix(makeappxPath, "does-not-exist", "")
		require.Errorf(t, err, "the system cannot find the directory: does-not-exist")
	})

	t.Run("package must be a file", func(t *testing.T) {
		err = packMsix("", "does-not-exist", "")
		require.Errorf(t, err, "the system cannot find the directory: does-not-exist")
	})

	fmt.Printf("Error: %s", err)
	err = os.RemoveAll(tmp)
	require.NoError(t, err)
}
