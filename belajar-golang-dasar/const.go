package main

import "fmt"

func main() {
	const firstName string = "Nasir"
	const lastName string = "Nooruddin"

	fmt.Println(firstName + " " + lastName)

	const(
		namaDepan string = "Nasir"
		namaBelakang string = "Nooruddin"
	)

	fmt.Println(namaDepan + " " + namaBelakang)
}