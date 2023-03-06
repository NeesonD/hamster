package main

import (
	"hamster/config"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Name string
}

func ReadFile(dirPaths []string) []*FileInfo {
	fileInfos := make([]*FileInfo, 0, 200)
	for _, dirPath := range dirPaths {
		// 遍历目录以及子目录下的所有文件
		filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			// 如果是文件夹则忽略
			if info.IsDir() {
				return nil
			}

			// 如果文件大小小于SkipDataSizeM则忽略
			if info.Size() < int64(config.Get().SkipDataSize)*1024*1024 {
				return nil
			}

			// 如果不是视频文件则忽略
			if isVideoFile(path) {
				return nil
			}

			// 记录文件的路径和名称
			fileInfos = append(fileInfos, &FileInfo{
				Path: path,
				Name: info.Name(),
			})

			return nil
		})
	}
	return fileInfos
}

// 判断文件是否是视频文件
func isVideoFile(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".mp4", ".mkv", ".avi", ".mov", ".wmv":
		return true
	default:
		return false
	}
}
