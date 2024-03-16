package main

import "fmt"

func main() {

	type NoKTP string

	var ktpNasir NoKTP = "11111111"

	fmt.Println(ktpNasir)
	fmt.Println(NoKTP("22222222"))

}
