package ziptools

import (
	"archive/zip"
	"fmt"
	"hash/crc32"
	"io"
)

func IsValidZip(file string) bool {
	// check if file is a valid zip file
	// if it is, return true
	// else return false
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		return false
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		if err := verifyZipFileDeep(file); err != nil {
			return false
		}
	}

	return true
}

func verifyZipFileDeep(file *zip.File) error {
	rc, err := file.Open()
	if err != nil {
		return fmt.Errorf("error opening file %s in zip: %w", file.Name, err)
	}
	defer rc.Close()

	h := crc32.NewIEEE()
	if _, err := io.Copy(h, rc); err != nil {
		return fmt.Errorf("error calculating CRC32 checksum for file %s: %w", file.Name, err)
	}

	if h.Sum32() != file.CRC32 {
		return fmt.Errorf("CRC32 checksum doesn't match for file %s", file.Name)
	}

	return nil
}
