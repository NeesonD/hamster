package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	fileIDMap     = map[string]interface{}{}
	failIdMap     = map[string]interface{}{}
	fileNameToUrl = map[string]interface{}{}
	fileIds       = "file_ids.json"
	failIds       = "fail_file_ids.json"
	nameUrl       = "fileNameUrl.json"
)

func main() {
	read(&fileIDMap, fileIds)
	read(&failIdMap, failIds)
	read(&fileNameToUrl, nameUrl)
	defer func() {
		write(fileIDMap, fileIds)
		write(failIdMap, failIds)
		write(fileNameToUrl, nameUrl)
	}()
	responses := parse()
	var wg sync.WaitGroup
	for _, response := range responses {
		for _, item := range response.Items {
			_, ok := fileIDMap[item.FileID]
			if !ok {
				err := createLink(item)
				if err != nil {
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	wg.Wait()
}

func createLink(item Item) error {
	url := "https://api.aliyundrive.com/adrive/v2/share_link/create"

	data := []byte(fmt.Sprintf(`{"drive_id":"18993090","expiration":"","share_pwd":"","share_name":"%s","file_id_list":["%s"]}`, item.Name, item.FileID))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("authority", "api.aliyundrive.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("authorization", token) // 将省略号替换为实际的令牌
	req.Header.Set("content-type", "application/json;charset-utf-8")
	req.Header.Set("cookie", "...")
	req.Header.Set("origin", "https://www.aliyundrive.com")
	req.Header.Set("referer", "https://www.aliyundrive.com/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="104"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) aDrive/4.1.0 Chrome/108.0.5359.215 Electron/22.3.1 Safari/537.36")
	req.Header.Set("x-canary", "client=windows,app=adrive,version=v4.1.0")
	req.Header.Set("x-device-id", "1dcb66bf-da6a-5443-b247-d184f198e839")
	req.Header.Set("x-request-id", "9bc6ec20-70a4-479e-9af3-f85385714ccb")
	req.Header.Set("x-signature", "9fe37f7a0995e7bdadb06f512203993c9051fa015c432933e91b3a15ebe5a2ef5df6ba518ad1d4a7e2928589f998a104c0d602490e2d13a1a5624c2848460a2201")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}
	var shareData ShareData
	err = json.Unmarshal(body, &shareData)
	if err != nil {
		fmt.Println("Error Unmarshal response body:", err)
		return err
	}
	if len(shareData.ShareURL) > 0 {
		fileIDMap[item.FileID] = shareData.ShareURL
		fileNameToUrl[item.Name] = shareData.ShareURL
		delete(failIdMap, item.FileID)
	} else {
		failIdMap[item.FileID] = item.Name
		fmt.Println("Error response body:", string(body))
		return err
	}

	fmt.Println("Response:", string(body))
	return nil
}

var (
	token = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI5OGQ4MmMxMWJmYWI0MWE5YjMzYmNlZTgwYzZkY2UzNyIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwicEpaSW5OSE4yZFpXazhxZ1wiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJGSUxFLkFMTFwiLFwiVklFVy5BTExcIixcIlNIQVJFLkFMTFwiLFwiU1RPUkFHRS5BTExcIixcIlNUT1JBR0VGSUxFLkxJU1RcIixcIlVTRVIuQUxMXCIsXCJCQVRDSFwiLFwiQUNDT1VOVC5BTExcIixcIklNQUdFLkFMTFwiLFwiSU5WSVRFLkFMTFwiLFwiU1lOQ01BUFBJTkcuTElTVFwiXSxcInJvbGVcIjpcInVzZXJcIixcInJlZlwiOlwiXCIsXCJkZXZpY2VfaWRcIjpcIjQwMDE2M2I2Zjc1ZDQ4YTJhODA3YjVlYmRkZDY0MTZlXCJ9IiwiZXhwIjoxNjgwNDE3MzU2LCJpYXQiOjE2ODA0MTAwOTZ9.RlvF5q0voLWhxZHuYl-MZbgUI1vpr6dLoUeyWcIbvY41jHSiIBcGlqn_TzWuiAMYfIjGlh-3ds_kOU4bcuIg7oXRUmUrnSqICX2knSknBeLo3jUdumvnJAbXZQh9LFN1JnbahjlkjyTxomvzX9MNd1g4-I8gfmz3YAdb6IW2VnM"
)

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
