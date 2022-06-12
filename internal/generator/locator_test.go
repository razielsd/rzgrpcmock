package generator

import (
	"github.com/stretchr/testify/require"
	"go/build"
	"os"
	"strings"
	"testing"
)

func Test_newPackageLocator(t *testing.T) {
	locator := newPackageLocator()
	require.NotNil(t, locator)
	sep := string(os.PathSeparator)
	modpath := strings.TrimSuffix(build.Default.GOPATH, sep) + sep + "pkg" + sep + "mod"
	require.Equal(t, modpath, locator.modpath)
	require.Equal(t, sep, locator.sep)
}

func Test_packageLocator_Search(t *testing.T) {
	locator := &packageLocator{modpath: "./data", sep: "/"}

	t.Run("success package found", func(t *testing.T) {
		path, err := locator.Search("my/package/here", "v1.3.1")
		require.NoError(t, err)
		require.Equal(t, "./data/my/package@v1.3.1/here", path)
	})

	t.Run("package not found", func(t *testing.T) {
		_, err := locator.Search("my/package/outside", "v1.3.1")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrPackageNotFound)
	})
}

func Test_packageLocator_makePath(t *testing.T) {
	locator := &packageLocator{modpath: "modpath", sep: "/"}

	t.Run("bad version index", func(t *testing.T) {
		path := locator.makePath("my/package/here", "v1.2.1", 3)
		require.Empty(t, path)
	})

	t.Run("last version index", func(t *testing.T) {
		path := locator.makePath("my/package/here", "v1.2.1", 0)
		require.Equal(t, "modpath/my/package/here@v1.2.1", path)
	})

	t.Run("middle version index", func(t *testing.T) {
		path := locator.makePath("my/package/here", "v1.2.1", 2)
		require.Equal(t, "modpath/my@v1.2.1/package/here", path)
	})
}
