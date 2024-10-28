package main

import "fmt"

func main() {
	var a, b int
	var hasil int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i += 2 {
		hasil = hasil + i
	}
	fmt.Print(hasil)
}
