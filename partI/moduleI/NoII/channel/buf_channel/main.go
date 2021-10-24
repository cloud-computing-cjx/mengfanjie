package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 6; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10) // 随机数范围: [0, 10)
			fmt.Println("putting: ", n)
			ch <- n
		}
		close(ch)
	}()
	fmt.Println("hello from main")
	for v := range ch {
		fmt.Println("receiving: ", v)
	}
}

/*
hello from main
putting:  8
putting:  2
putting:  8
putting:  3
putting:  1
putting:  3
receiving:  8
receiving:  2
receiving:  8
receiving:  3
receiving:  1
receiving:  3
*/
