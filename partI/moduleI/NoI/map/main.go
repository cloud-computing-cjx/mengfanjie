package main

import "fmt"

func main() {
	// map key:string value:string
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"

	// map key:string value:func
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())

	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}

/*
➜  NoI git:(main) ✗ go run map/main.go
map[funcA:0x47ecc0]
1
b
a b
*/
