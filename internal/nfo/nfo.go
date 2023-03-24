package nfo

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var outlineReg = regexp.MustCompile(`<tagline>(.*?)</tagline>`)

func AppendLink(filePath string, link string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 使用正则表达式匹配<outline>和</outline>之间的文本
	result := outlineReg.ReplaceAllString(string(file), fmt.Sprintf(`<tagline>$1 link:%s</tagline>`, link))

	err = ioutil.WriteFile(filePath, []byte(result), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("xue.xml 文件已修改成功！")
}
