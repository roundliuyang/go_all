package demo

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

// 对于大规模的数据处理，excelize 提供了流式读写 API，避免内存占用过高。
// 使用 StreamWriter 实现大规模数据的流式写入：
func TestFlush(t *testing.T) {
	// 创建新的 Excel 文件
	f := excelize.NewFile()

	// 获取 StreamWriter 对象（默认新建文件中 Sheet1 存在）
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Fatalf("创建 StreamWriter 失败: %v", err)
	}

	// 写入一百万行数据
	for row := 1; row <= 1000; row++ {
		cell, _ := excelize.CoordinatesToCellName(1, row) // A列，第 row 行
		err := streamWriter.SetRow(cell, []interface{}{fmt.Sprintf("Row %d", row)})
		if err != nil {
			log.Fatalf("写入第 %d 行失败: %v", row, err)
		}
	}

	// 刷新缓冲数据
	if err := streamWriter.Flush(); err != nil {
		log.Fatalf("刷新写入数据失败: %v", err)
	}

	// 保存文件
	if err := f.SaveAs("LargeData.xlsx"); err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}

	fmt.Println("成功写入 1000 行数据至 LargeData.xlsx")
}
