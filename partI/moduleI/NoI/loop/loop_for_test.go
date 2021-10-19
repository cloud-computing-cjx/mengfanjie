package main

import "testing"

// Go 只有一种循环结构：for 循环
func Testfor(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	t.Log(sum)
}
