package sort

import (
	"testing"
)

// go test -v -run=TestMergeSort
func TestMergeSort(t *testing.T) {
	var nums []int = []int{11, 2, 66, 9, 12, 3, 6, 9}
	t.Log(nums)
	t.Log(MergeSort(nums))
	// 测试一个map 引用问题 和 排序无关下面代码
	var maps map[int]string
	maps = make(map[int]string)
	c_map := make(map[int]string)
	maps[1] = "Hello"
	t.Log(maps)
	c_map = maps
	c_map[2] = "World"
	t.Log(maps[2])
	c_map = nil
	t.Log(maps[2])
	t.Log(c_map == nil)

}
