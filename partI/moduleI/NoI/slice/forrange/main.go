package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	// Go 中默认都是值传递
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

	// 修改切片的值
	for index := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}

/*
➜  forrange git:(main) ✗ go run main.go
mySlice [1 2 3 4 5]
mySlice [2 4 6 8 10]
*/
