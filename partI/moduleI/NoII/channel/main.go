package main

import "fmt"

// 如果不使用 channel, 就需要使用 sleep 的方式保证 goroutine 的结束
func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0 //数据写入 Channel
	}()
	i := <-ch // 从 Channel 中取数据并赋值
	fmt.Printf("%d", i)
}

/*
hello from goroutine
0
*/
