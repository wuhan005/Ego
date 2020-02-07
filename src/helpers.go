package main

import (
	"errors"
	"io/ioutil"
	"os"
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
