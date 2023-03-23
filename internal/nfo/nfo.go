package nfo

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
		panic(err)
	}

	fmt.Printf("%s 文件已修改成功！\n", file)
}
