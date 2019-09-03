package gopack

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rogpeppe/go-internal/modfile"
)

// GetModuleNameFromModfile reads a go.mod file from the given directory and
// parses out the module name before returning it to the caller.
func GetModuleNameFromModfile(dir string) (string, error) {
	content, err := loadModfileFromDisk(dir)
	if err != nil {
		return "", err
	}
	return parseModuleNameFromModfile(content)
}

func loadModfileFromDisk(dir string) ([]byte, error) {
	modFilePath := filepath.Join(dir, "go.mod")
	modFile, err := os.Open(modFilePath) //nolint:gosec
	if err != nil {
		return []byte{}, err
	}
	defer modFile.Close()

	return ioutil.ReadAll(modFile)
}

func parseModuleNameFromModfile(modfileContent []byte) (string, error) {
	modulePath := modfile.ModulePath(modfileContent)
	if modulePath == "" {
		return "", errors.New("unable to parse module path from go.mod contents")
	}
	return modulePath, nil
}
