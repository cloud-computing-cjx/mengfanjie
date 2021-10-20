package main

func main() {
	DoOperation2(1, increase)
	DoOperation(1, decrease)
}

func increase(a, b int) int {
	return a + b
}

func DoOperation2(y int, f func(int, int) int) {
	println("increase result is:", f(y, 1))
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
