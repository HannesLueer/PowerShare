package helper

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// CreateDir creates the specified directory with permissions 0755
func CreateDir(path string) error{
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			return errDir
		}
	}

	return nil
}

// Exists returns if the specified directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// IsEmpty check is the given directory is empty
func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// WriteFile creates a file at the given path and writes the data to it
func WriteFile(data []byte, path string) error {
	CreateDir(filepath.Dir(filepath.Dir(path)))

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		file.Close()
		return err
	}

	return nil
}
