package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 5
	var d = 5
	var f = a + b - c*d

	fmt.Println(f)

	var i = 10
	i += 10

	fmt.Println(i)

	i += 5

	fmt.Println(i)
}
