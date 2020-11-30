package utils

import (
	"math/rand"
	"time"
)

func RandNums(limit int) (nums []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < limit; i++ {
		nums = append(nums, (rand.Intn(99999) - 88888))
	}
	return
}
