package generator

import (
	"errors"
	"os"
	"strings"
)

var ErrPackageNotFound = errors.New("package not found")

type packageLocator struct {
	PathList []string
	modpath  string
	sep      string
}

func newPackageLocator(projectDir string) *packageLocator {
	sep := string(os.PathSeparator)
	return &packageLocator{
		sep:     sep,
		modpath: strings.TrimSuffix(projectDir, string(os.PathSeparator)) + sep + "vendor",
	}
}

func (p *packageLocator) Search(name string) (string, error) {
	path := p.modpath + p.sep + name
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	return "", ErrPackageNotFound
}
