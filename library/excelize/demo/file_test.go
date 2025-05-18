package demo

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

// 创建一个名为 Sheet1 的默认工作表，并在单元格 A1 和 B1 中写入了数据。随后将 Excel 文件保存为 Book1.xlsx
func TestCreateFile(t *testing.T) {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 设置单元格的值
	f.SetCellValue("Sheet1", "A1", "Hello, Excelize!")
	f.SetCellValue("Sheet1", "B1", 100)

	// 保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// 使用 excelize.OpenFile() 方法打开现有的 Excel 文件，并读取单元格中的内容
func TestReadeFile(t *testing.T) {
	// 打开现有的 Excel 文件
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取单元格的值
	cellValue, err := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	cellValue2, err := f.GetCellValue("Sheet1", "B1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("A1 Cell Value:", cellValue)
	fmt.Println("B1 Cell Value:", cellValue2)

	// 关闭文件
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
