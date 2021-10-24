package main

import "fmt"

// Go 中只有两种结构: struct 和 interface. struct 只包含属性, interface 只包含方法
type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.firstName + ", " + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())
}

/*
➜  mengfanjie git:(main) go run partI/moduleI/NoII/interface/main.go
first, last
benz-s
vendor: testVendor, model: testModel
*/
