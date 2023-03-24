package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	videoExtensions  = ".avi,.mp4,.mov"  // 视频格式后缀名
	torrentExtension = ".torrent"        // torrent 文件后缀名
	maxFileSize      = 120 * 1024 * 1024 // 文件最大大小（120M）
)

func DeleteFiles(dirPath string) error {
	// 遍历目录下所有文件和子目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			// 如果是子目录，则递归删除
			err := DeleteFiles(filePath)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，则判断是否需要删除
			if shouldDeleteFile(filePath) {
				err := os.Remove(filePath)
				if err != nil {
					return err
				}
				fmt.Printf("Deleted file: %s\n", filePath)
			}
		}
	}
	return nil
}

func shouldDeleteFile(filePath string) bool {
	// 判断是否需要删除文件
	extension := strings.ToLower(filepath.Ext(filePath))
	fileSize := getFileSize(filePath)
	return (strings.Contains(videoExtensions, extension) ||
		extension == torrentExtension) && fileSize < maxFileSize
}

func getFileSize(filePath string) int64 {
	// 获取文件大小
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
