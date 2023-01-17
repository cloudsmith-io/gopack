package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileFilter = func(path string) bool

// CreateModuleArchive zips up relevant files to create a compliant Go module zip file.
func CreateModuleArchive(dir, modPath, version string, filter FileFilter) error {
	archiveFileName := fmt.Sprintf("%s.zip", version)
	writer, err := os.Create(filepath.Join(dir, archiveFileName))
	if err != nil {
		return err
	}
	defer writer.Close()

	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil || info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if strings.Contains(path, ".git") || strings.HasSuffix(path, archiveFileName) {
			return nil
		}
		if filter != nil && filter(path) {
			return nil
		}

		fileName := generateFileName(dir, path, modPath, version)
		file, err := os.Open(path) //nolint:gosec
		if err != nil {
			return err
		}
		defer file.Close()

		zipFile, err := zipWriter.Create(fileName)
		if err != nil {
			return err
		}

		_, err = io.CopyN(zipFile, file, info.Size())
		return err
	})
}

func generateFileName(dir, filePath, modPath, version string) string {
	fileName := strings.TrimPrefix(filePath, dir)
	fileName = strings.TrimLeft(fileName, string(os.PathSeparator))

	return filepath.Join(fmt.Sprintf("%s@%s", modPath, version), fileName)
}
