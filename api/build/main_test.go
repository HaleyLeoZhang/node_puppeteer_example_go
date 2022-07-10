package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func add(a, b int) int {
	return a + b
}

func do() {
	var c int
	once.Do(func() {
		c = add(1, 2)
		<-time.After(1 * time.Second) // 让结果延迟，展示得更明显
	})
	fmt.Println("c的值", c)
}

func TestService_Add(t *testing.T) {
	go do()
	go do()
	<-time.After(3 * time.Second) // 让结果延迟，展示得更明显
}
