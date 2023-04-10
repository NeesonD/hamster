package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	TagLine = iota
	Plot
)

type Tag struct {
	Reg    string
	Format string
}

var (
	tagMap = map[int]Tag{
		TagLine: {
			Reg:    "<tagline>(.*?)</tagline>",
			Format: "<tagline>%s</tagline>",
		},
		Plot: {
			Reg:    "<plot>(.*?)</plot>",
			Format: "<plot>$1</plot>\n  <tagline>%s</tagline>",
		},
	}
	taglineReg = regexp.MustCompile(tagMap[TagLine].Reg)
	plotReg    = regexp.MustCompile(tagMap[Plot].Reg)
)

var (
	fileNameUrl = "fileNameUrl.json"
	srcFilePath = "Z:\\阿里云盘Open\\4K\\4K原盘"
	fileIDMap   = map[string]interface{}{}
	failIdMap   = map[string]interface{}{}
	fileIds     = "newfile_ids.json"
	failIds     = "newfail_file_ids.json"
)

func main() {
	read(&fileIDMap, fileIds)
	read(&failIdMap, failIds)
	defer func() {
		write(fileIDMap, fileIds)
		write(failIdMap, failIds)
	}()
	type JsonData map[string]string
	// 读取 JSON 文件
	jsonFileData, err := ioutil.ReadFile(fileNameUrl)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var jsonData JsonData
	err = json.Unmarshal(jsonFileData, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	err = filepath.Walk(srcFilePath, func(path string, info os.FileInfo, err error) error {
		fmt.Println("curr path", path)
		defer func() {
			if i := recover(); i != nil {
				fmt.Println(i)
			}
		}()
		if !info.IsDir() && filepath.Ext(path) == ".nfo" {
			split := strings.Split(path, "\\")
			p := split[len(split)-2]
			AppendLink(path, "")
			fileIDMap[p] = ""
			delete(failIdMap, p)
		}
		return nil
	})
	if err != nil {
		fmt.Println("失败", err)
	}
}

func AppendLink(filePath string, link string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fs := string(file)
	result := ""
	match := taglineReg.FindStringSubmatch(fs)
	if len(match) > 1 {
		result = taglineReg.ReplaceAllString(fs, fmt.Sprintf(tagMap[TagLine].Format, link))
	} else {
		result = plotReg.ReplaceAllString(fs, fmt.Sprintf(tagMap[Plot].Format, link))
	}

	err = ioutil.WriteFile(filePath, []byte(result), 0644)
	if err != nil {
		fmt.Println("nfo err", err)
	}

}

func write(data interface{}, file string) {
	// 1. 将 fileId 数组存储到json文件
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile(file, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func read(data interface{}, file string) {

	// 2. 读取json文件
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

}
