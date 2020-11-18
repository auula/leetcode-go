package search

// BinarySearch 对查找的数据进行折半查找 降低数据量
// https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%88%86%E6%90%9C%E5%B0%8B%E6%BC%94%E7%AE%97%E6%B3%95#golang_%E9%80%92%E5%BD%92%E7%89%88%E6%9C%AC
func BinarySearch(nums []int, key int) (index int) {
	var middle, left, right int = 0, 0, len(nums) - 1
	for left <= right {
		middle = left + (right-left)/2
		if key == nums[middle] {
			index = middle
		}
		if key < nums[middle] {
			// 缩小范围
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return index
}
