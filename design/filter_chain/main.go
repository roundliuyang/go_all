package main

import (
	"fmt"
	"log"
)

type MediaFilter interface {
	Preflight() bool
	Filter() ([]int, error)
}

type CopyrightMediaFilter struct {
	// 依赖注入的内容
	classID int
}

func (f *CopyrightMediaFilter) Preflight() bool {
	return f.classID != 0
}

func (f *CopyrightMediaFilter) Filter() ([]int, error) {
	// 模拟返回
	return []int{1, 2, 3, 4}, nil
}

type SyncStatusMediaFilter struct {
	status string
}

func (f *SyncStatusMediaFilter) Preflight() bool {
	return f.status != ""
}

func (f *SyncStatusMediaFilter) Filter() ([]int, error) {
	return []int{2, 3}, nil
}

type FuzzyNameMediaFilter struct {
	name string
}

func (f *FuzzyNameMediaFilter) Preflight() bool {
	return f.name != ""
}

func (f *FuzzyNameMediaFilter) Filter() ([]int, error) {
	return []int{3, 4, 5}, nil
}

// 组合执行逻辑（核心部分）
func FilterMedia(filters []MediaFilter) ([]int, error) {
	var filterIds []int

	for _, filter := range filters {
		if !filter.Preflight() {
			continue
		}

		list, err := filter.Filter()
		if err != nil {
			return nil, err
		}
		if list == nil {
			return nil, nil
		}

		if filterIds == nil {
			filterIds = list
		} else {
			filterIds = intersection(filterIds, list)
		}
	}

	return filterIds, nil
}

// 求交集
func intersection(a, b []int) []int {
	set := make(map[int]bool)
	for _, v := range a {
		set[v] = true
	}
	var result []int
	for _, v := range b {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}

// 执行示例
func main() {
	filters := []MediaFilter{
		&CopyrightMediaFilter{classID: 123},
		&SyncStatusMediaFilter{status: "SYNCED"},
		&FuzzyNameMediaFilter{name: "sample"},
	}

	filteredIds, err := FilterMedia(filters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Filtered IDs:", filteredIds)
}
