package demo

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

// 在 Excel 文件中，你可以轻松地操作单元格内容。excelize 提供了丰富的 API 用于写入、读取和修改单元格数据。

// 使用 SetCellValue() 可以向指定单元格写入数据。可以写入多种类型的数据，包括字符串、数字、布尔值等
// 你可以使用 GetCellValue() 方法读取单元格的数据，返回的数据总是字符串类型：
func TestSetAndGetCellValue(t *testing.T) {
	// 创建新文件
	f := excelize.NewFile()

	// 向 Sheet1 写入不同类型的数据
	f.SetCellValue("Sheet1", "A1", "Go语言")
	f.SetCellValue("Sheet1", "B1", 12345)
	f.SetCellValue("Sheet1", "C1", true)

	// 保存文件
	if err := f.SaveAs("TestCellValue.xlsx"); err != nil {
		t.Fatalf("保存文件失败: %v", err)
	}

	// 重新打开文件进行读取验证
	f2, err := excelize.OpenFile("TestCellValue.xlsx")
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}

	// 读取并验证 A1
	if value, err := f2.GetCellValue("Sheet1", "A1"); err != nil {
		t.Errorf("读取 A1 失败: %v", err)
	} else if value != "Go语言" {
		t.Errorf("A1 期望值: Go语言，实际值: %s", value)
	}

	// 读取并验证 B1
	if value, err := f2.GetCellValue("Sheet1", "B1"); err != nil {
		t.Errorf("读取 B1 失败: %v", err)
	} else if value != "12345" {
		t.Errorf("B1 期望值: 12345，实际值: %s", value)
	}

	// 读取并验证 C1
	if value, err := f2.GetCellValue("Sheet1", "C1"); err != nil {
		t.Errorf("读取 C1 失败: %v", err)
	} else if value != "TRUE" {
		t.Errorf("C1 期望值: TRUE，实际值: %s", value)
	}
}

func TestGetCellTypeFromExistingFile(t *testing.T) {
	// 打开已有文件
	f, err := excelize.OpenFile("TestCellValue.xlsx")
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}

	cellType, err := f.GetCellType("Sheet1", "C1")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cellType {
	case excelize.CellTypeNumber:
		fmt.Println("A1 contains a number.")
	case excelize.CellTypeInlineString, excelize.CellTypeSharedString:
		fmt.Println("A1 contains a string.")
	case excelize.CellTypeBool:
		fmt.Println("A1 contains a boolean.")
	case excelize.CellTypeUnset:
		fmt.Println("A1 contains a CellTypeUnset.")
	default:
		fmt.Printf("A1 contains type: %v\n", cellType)
	}
}

// 单元格样式
func TestSetCellBorder(t *testing.T) {
	// 打开已有文件
	f, err := excelize.OpenFile("TestCellValue.xlsx")
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}

	// 新建边框样式，左右边框黑色，细线（style=1）
	borderStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
	})
	if err != nil {
		t.Fatalf("创建边框样式失败: %v", err)
	}

	// 设置单元格 A1 的样式
	if err := f.SetCellStyle("Sheet1", "B1", "B1", borderStyle); err != nil {
		t.Fatalf("设置单元格样式失败: %v", err)
	}

	// 保存为新文件，防止覆盖原文件
	saveFile := "TestCellValue_withBorder.xlsx"
	if err := f.SaveAs(saveFile); err != nil {
		t.Fatalf("保存文件失败: %v", err)
	}

	t.Logf("成功为单元格 A1 设置边框，并保存为 %s", saveFile)
}

func TestMergeCellA1B1(t *testing.T) {
	// 打开已有 Excel 文件
	f, err := excelize.OpenFile("TestCellValue.xlsx")
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}

	// 合并 Sheet1 的 A1 和 B1 单元格
	if err := f.MergeCell("Sheet1", "A1", "A2"); err != nil {
		t.Fatalf("合并单元格失败: %v", err)
	}

	// 保存为新文件，防止覆盖原始文件
	saveAs := "TestCellValue_Merged.xlsx"
	if err := f.SaveAs(saveAs); err != nil {
		t.Fatalf("保存文件失败: %v", err)
	}

	t.Logf("成功合并 Sheet1 的 A1 和 B1，并保存为 %s", saveAs)
}
