package olapdb

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getClientArea(t *testing.T) {
	p, _ := filepath.Abs(filepath.Join("./testdata", "GeoIP2-City-Test.mmdb"))
	err := OpenGeoDB(p)
	require.NoError(t, err)
	assert.Equal(t, "gb", getClientArea("81.2.69.142"))
	assert.Equal(t, "", getClientArea("2001:41d0:303:df3e::"))
}
