package generator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newPackageLocator(t *testing.T) {
	locator := newPackageLocator(".")
	require.NotNil(t, locator)
	sep := string(os.PathSeparator)
	modpath := "./vendor"
	require.Equal(t, modpath, locator.modpath)
	require.Equal(t, sep, locator.sep)
}

func Test_packageLocator_Search(t *testing.T) {
	locator := &packageLocator{modpath: "./data", sep: "/"}

	t.Run("success package found", func(t *testing.T) {
		path, err := locator.Search("my/package/here")
		require.NoError(t, err)
		require.Equal(t, "./data/my/package/here", path)
	})

	t.Run("package not found", func(t *testing.T) {
		_, err := locator.Search("my/package/outside")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrPackageNotFound)
	})
}
