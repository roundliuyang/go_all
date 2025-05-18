package demo

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

// 工作表操作, excelize 支持创建、删除和操作多个工作表。你可以为每个 Excel 文件创建多个工作表，并在不同工作表之间切换

// 创建新工作表
func TestCreateSheet(t *testing.T) {
	f := excelize.NewFile()
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		panic(err)
	}
	f.SetCellValue("Sheet2", "A1", "Hello from Sheet2")
	f.SetActiveSheet(index)
	f.SaveAs("Book2.xlsx")
}

// 向已有的 Book2.xlsx 添加 Sheet2 工作表
func TestAddSheet(t *testing.T) {
	// 打开已有文件
	f, err := excelize.OpenFile("Book2.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	// 如果 Sheet2 已存在，跳过创建
	if index, _ := f.GetSheetIndex("Sheet2"); index != -1 {
		t.Log("Sheet2 已存在，跳过创建")
		return
	}

	// 创建 Sheet2
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		t.Fatal(err)
	}

	// 设置内容
	f.SetCellValue("Sheet2", "A1", "Hello from newly added Sheet2")

	// 设置为活动工作表（可选）
	f.SetActiveSheet(index)

	// 保存更改
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
}

// 删除工作表
func TestDeleteSheet(t *testing.T) {
	f, err := excelize.OpenFile("Book2.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	// 确保当前激活工作表不是 Sheet2
	if f.GetSheetName(f.GetActiveSheetIndex()) == "Sheet2" {
		// 切换到 Sheet1 或其他工作表
		index, _ := f.GetSheetIndex("Sheet1")
		if index == -1 {
			t.Fatal("Sheet1 不存在，无法设置为激活工作表")
		}
		f.SetActiveSheet(index)
	}

	// 删除 Sheet2
	if err := f.DeleteSheet("Sheet2"); err != nil {
		t.Fatal(err)
	}

	// 保存修改
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
}

// 重命名工作表
func TestRenameSheet(t *testing.T) {
	// 打开已有文件
	f, err := excelize.OpenFile("Book2.xlsx")
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}

	// 检查 "Sheet1" 是否存在
	if index, _ := f.GetSheetIndex("Sheet1"); index == -1 {
		t.Fatalf("Sheet1 不存在，无法重命名")
	}

	// 重命名 Sheet1 为 NewSheetName
	if err := f.SetSheetName("Sheet1", "NewSheetName"); err != nil {
		t.Fatalf("重命名工作表失败: %v", err)
	}

	// 保存修改
	if err := f.Save(); err != nil {
		t.Fatalf("保存文件失败: %v", err)
	}

	t.Log("工作表 Sheet1 已成功重命名为 NewSheetName")
}
