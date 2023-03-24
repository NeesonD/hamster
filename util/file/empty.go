package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EmptyDirs(path string) ([]string, error) {
	var emptyDirs []string

	err := filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && subpath != path {
			isEmpty, _ := IsDirEmpty(subpath)
			if isEmpty {
				emptyDirs = append(emptyDirs, subpath)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return emptyDirs, nil
}

func IsDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func FindEmpty(dir string) {
	emptyDirs, err := EmptyDirs(dir)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Empty directories:\n")
	for _, dir := range emptyDirs {
		fmt.Printf("- %s\n", dir)
	}
}
