package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// defer 以栈的方式执行
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// 匿名函数
		go func(i int) {
			lock.Lock()
			// defer 在方法结束时执行
			defer lock.Unlock()
			fmt.Println("loopFunc: ", i)
		}(i)
	}
}

/*
loopFunc:  2
loopFunc:  0
loopFunc:  1
3
2
1
*/
