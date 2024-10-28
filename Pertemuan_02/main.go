package main

import "fmt"

func main() {
	var a1, a2, a3 float32
	var hasil1, hasil2, hasil3 float32

	//input
	fmt.Scan(&a1, &a2, &a3)

	//rumus menghitung keuntungan 5%
	hasil1 = a1 + (a1 * (5.0 / 100))
	hasil2 = a2 + (a2 * (5.0 / 100))
	hasil3 = a3 + (a3 * (5.0 / 100))

	//output
	fmt.Println(hasil1, hasil2, hasil3)
}
