package main

import (
	"fmt"

	_ "github.com/cncamp/mengfanjie/partI/moduleI/NoII/init/a"
	_ "github.com/cncamp/mengfanjie/partI/moduleI/NoII/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("main")
}
