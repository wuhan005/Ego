package main

import (
	"errors"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func isExist(path string, isDir bool) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}
	if isDir && !file.IsDir() {
		return errors.New(path + "should be a folder")
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func CopyFile(fs afero.Fs, from, to string) error {
	sf, err := os.Open(from)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(to)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(from)
		if err != nil {
			err = os.Chmod(to, si.Mode())

			if err != nil {
				return err
			}
		}

	}
	return nil
}

// CopyDir copies a directory.
func CopyDir(fs afero.Fs, from, to string, shouldCopy func(filename string) bool) error {
	fi, err := os.Stat(from)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return errors.New(from + " is not a directory")
	}

	err = fs.MkdirAll(to, 0777) // before umask
	if err != nil {
		return err
	}

	entries, _ := ioutil.ReadDir(from)
	for _, entry := range entries {
		fromFilename := filepath.Join(from, entry.Name())
		toFilename := filepath.Join(to, entry.Name())
		if entry.IsDir() {
			if shouldCopy != nil && !shouldCopy(fromFilename) {
				continue
			}
			if err := CopyDir(fs, fromFilename, toFilename, shouldCopy); err != nil {
				return err
			}
		} else {
			if err := CopyFile(fs, fromFilename, toFilename); err != nil {
				return err
			}
		}

	}

	return nil
}
