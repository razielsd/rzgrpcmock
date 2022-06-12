package generator

import (
	"errors"
	"go/build"
	"os"
	"strings"
)

var ErrPackageNotFound = errors.New("package not found")

type packageLocator struct {
	PathList []string
	modpath  string
	sep      string
}

func newPackageLocator() *packageLocator {
	sep := string(os.PathSeparator)
	return &packageLocator{
		sep:     sep,
		modpath: strings.TrimSuffix(build.Default.GOPATH, string(os.PathSeparator)) + sep + "pkg" + sep + "mod",
	}
}

func (p *packageLocator) Search(name, version string) (string, error) {
	index := 0
	for {
		path := p.makePath(name, version, index)
		index++
		if path == "" {
			break
		}
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", ErrPackageNotFound
}

func (p *packageLocator) makePath(name, version string, index int) string {
	parts := strings.Split(name, p.sep)
	if len(parts) <= index {
		return ""
	}
	parts[len(parts)-index-1] += "@" + version

	return p.modpath + p.sep + strings.Join(parts, p.sep)
}
