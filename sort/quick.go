package sort

// QuickSort 从无序数据里面取到一个中间值 然后比较 大的放到这个值的右边 小的放到左边
// https://mojotv.cn/algorithm/golang-quick-sort
func QuickSort(nums []int, low, hight int) {
	if low >= hight {
		return
	}
	pivot := nums[low]
	L, R := low, hight
	for low < hight {
		for low < hight && nums[hight] >= pivot {
			hight--
		}
		nums[low] = nums[hight]
		for low < hight && nums[low] <= pivot {
			low++
		}
		nums[hight] = nums[low]
	}
	nums[low] = pivot
	QuickSort(nums, L, low-1)
	QuickSort(nums, low+1, R)
}
