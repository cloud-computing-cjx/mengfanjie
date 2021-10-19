package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// name 的默认值设为 world, (key, default_value, desc)
	name := flag.String("name", "world", "specify the name you want to say hi")
	// 解析 os.Args，将 --key 后的值赋值给 key
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
