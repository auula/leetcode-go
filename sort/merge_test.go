package sort

import (
	"testing"

	"github.com/higker/leetcode-go/utils"
)

// go test -v -run=TestMergeSort
// === RUN   TestMergeSort
// --- PASS: TestMergeSort (43.84s)
// PASS
// ok      github.com/higker/leetcode-go/sort      44.156s
func TestMergeSort(t *testing.T) {
	var nums []int = utils.RandNums(100000000)
	//t.Log(nums)
	//t.Log(MergeSort(nums))
	MergeSort(nums)
	// 测试一个map 引用问题 和 排序无关下面代
	// var maps map[int]string
	// maps = make(map[int]string)
	// c_map := make(map[int]string)
	// maps[1] = "Hello"
	// t.Log(maps)
	// c_map = maps
	// c_map[2] = "World"
	// t.Log(maps[2])
	// c_map = nil
	// t.Log(maps[2])
	// t.Log(c_map == nil)
	// t.Log(maps == nil)

}
