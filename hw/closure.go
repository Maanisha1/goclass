package main

import "fmt"

func main() {
	m := printInt()

	fmt.Println(m)
	fmt.Println(m)

}
func printInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
