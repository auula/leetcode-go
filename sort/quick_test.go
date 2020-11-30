package sort

import (
	"testing"

	"github.com/higker/leetcode-go/utils"
)

// go test -v -run=TestQuickSort
// === RUN   TestQuickSort
// --- PASS: TestQuickSort (35.00s)
// PASS
// ok      github.com/higker/leetcode-go/sort      35.085s
func TestQuickSort(t *testing.T) {
	var nums []int = utils.RandNums(100000000)
	QuickSort(nums, 0, len(nums)-1)
}
