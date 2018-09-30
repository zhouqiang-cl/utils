package utils

import (
	"math/rand"

	"github.com/ngaut/log"
)

// GetRandomN gets random num of int, boundary is [0, total)
func GetRandomN(total, num int) []int {
	if num > total {
		log.Warnf("the num %d is greater than total %d", num, total)
		num = total
	}
 	totalArray := make([]int, 0, total)
	for i := 0; i < total; i++ {
		totalArray = append(totalArray, i)
	}
 	for j := 0; j < num; j++ {
		r := j + rand.Intn(total-j)
		totalArray[j], totalArray[r] = totalArray[r], totalArray[j]
	}
 	return totalArray[:num]
}