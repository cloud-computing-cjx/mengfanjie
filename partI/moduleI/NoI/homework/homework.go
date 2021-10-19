package main

import (
	"fmt"
)

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	for idx, _ := range myArray {
		if idx == 2 {
			myArray[idx] = "smart"
		}
		if idx == 4 {
			myArray[idx] = "strong"
		}
	}

	for _, value := range myArray {
		fmt.Printf("%s ", value)
	}
}
