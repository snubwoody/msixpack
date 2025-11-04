package bundle

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	temp := os.TempDir()
	dest := path.Join(temp, "test-pdf")
	url := "https://file-examples.com/wp-content/storage/2017/10/file-sample_150kB.pdf"
	p, err := downloadFile(url, dest)
	require.NoError(t, err)

	_, err = os.Stat(p)
	require.NoError(t, err)
}
