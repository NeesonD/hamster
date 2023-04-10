package file

import (
	"os"
	"path/filepath"
)

func DirSize(path string) int64 {
	var size int64

	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size
}

func GeSize(m int64, size int64) bool {
	return size > m*1024*1024
}
