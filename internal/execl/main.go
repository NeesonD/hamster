package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"sort"
)

type JsonData map[string]string

var (
	fileNameUrl  = "fileNameUrl.json"
	csvFilePath  = "./原盘电影-1065部-56TB-刮削.csv"
	xlsxFilePath = "./原盘电影-1065部-56TB-刮削.xlsx"
)

func main() {
	noNo()
}

func hasNo() {
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

	// 按 key 排序
	keys := make([]string, 0, len(jsonData))
	for k := range jsonData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建 Excel 文件
	excelFile := excelize.NewFile()
	sheet := "Sheet1"
	excelFile.SetSheetName("Sheet1", sheet)

	// 设置 Excel 表头
	excelFile.SetCellValue(sheet, "A1", "编号")
	excelFile.SetCellValue(sheet, "B1", "名称")
	excelFile.SetCellValue(sheet, "C1", "分享链接")

	// 创建 CSV 文件
	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	_ = csvWriter.Write([]string{"编号", "名称", "分享链接"}) // CSV 表头

	// 将排序后的数据写入 Excel 和 CSV
	for index, key := range keys {
		value := jsonData[key]

		// 写入 Excel
		excelFile.SetCellValue(sheet, fmt.Sprintf("A%d", index+2), index+1)
		excelFile.SetCellValue(sheet, fmt.Sprintf("B%d", index+2), key)
		excelFile.SetCellValue(sheet, fmt.Sprintf("C%d", index+2), value)

		// 写入 CSV
		_ = csvWriter.Write([]string{fmt.Sprint(index + 1), key, value})
	}

	// 保存 Excel 文件
	err = excelFile.SaveAs(xlsxFilePath)
	if err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	// 保存 CSV 文件
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		fmt.Println("Error saving CSV file:", err)
		return
	}

	fmt.Println("JSON data has been successfully written to Excel and CSV files.")
}
func noNo() {
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

	// 按 key 排序
	keys := make([]string, 0, len(jsonData))
	for k := range jsonData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建 Excel 文件
	excelFile := excelize.NewFile()
	sheet := "Sheet1"
	excelFile.SetSheetName("Sheet1", sheet)

	// 设置 Excel 表头
	excelFile.SetCellValue(sheet, "A1", "名称")
	excelFile.SetCellValue(sheet, "B1", "分享链接")

	// 创建 CSV 文件
	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	_ = csvWriter.Write([]string{"名称", "分享链接"}) // CSV 表头

	// 将排序后的数据写入 Excel 和 CSV
	for index, key := range keys {
		value := jsonData[key]

		// 写入 Excel
		excelFile.SetCellValue(sheet, fmt.Sprintf("A%d", index+2), key)
		excelFile.SetCellValue(sheet, fmt.Sprintf("B%d", index+2), value)

		// 写入 CSV
		_ = csvWriter.Write([]string{key, value})
	}

	// 保存 Excel 文件
	err = excelFile.SaveAs(xlsxFilePath)
	if err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	// 保存 CSV 文件
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		fmt.Println("Error saving CSV file:", err)
		return
	}

	fmt.Println("JSON data has been successfully written to Excel and CSV files.")
}
