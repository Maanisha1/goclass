package main

import "fmt"

func main() {
	a := printArea()
	m := a(2, 4)
	fmt.Println("result :", m)

}

func printArea() func(int, int) {
	return func(a, b int) int {
		return a * b
	}

}
