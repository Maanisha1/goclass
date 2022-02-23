package main

import "fmt"

func main() {
	m := printSum(2, 3)
	fmt.Println("m is :", m)

}

func printSum(a int, b int) int {
	Result := a + b
	return int(Result)

}
