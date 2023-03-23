package nfo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppendLink(t *testing.T) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".nfo" {
			AppendLink(path, "http://aliyun2")
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
