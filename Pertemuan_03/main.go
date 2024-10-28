package main

import "fmt"

func main() {
	var bilangan int
	var d1, d2, d3 int

	//input
	fmt.Scan(&bilangan)

	//pisahkan setiap digit dengan menggunakan modulus
	d1 = bilangan / 100
	d2 = bilangan / 10 % 10
	d3 = bilangan % 10

	//output
	fmt.Println(d1 < d2 && d2 < d3 || d1 > d2 && d2 > d3)
}
