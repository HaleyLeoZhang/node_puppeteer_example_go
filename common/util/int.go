package util

import (
	"math/rand"
	"time"
)

// 获取区间内随机数

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
