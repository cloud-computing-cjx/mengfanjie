package main

import (
	"fmt"
	"time"
)

func main() {
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	for i := 0; i < 6; i++ {
		go fmt.Println(i)
	}
}

/*
5
2
3
4
0
1
*/
