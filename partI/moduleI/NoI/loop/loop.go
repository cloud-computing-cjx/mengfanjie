package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println()
	fmt.Printf("%d", sum)

	fullString := "hello world"
	fmt.Println(fullString)
	for index, char := range fullString {
		fmt.Println(index, string(char))
	}
}
