package main

import "fmt"

func main() {
	var i, N, d int
	var s bool

	//input
	fmt.Scan(&N)

	//looping
	for i = 1; i <= N; i++ {

		d = i //mengambil nilai dari i
		//jika N hasil bagi i == 0 maka itu adalah faktor nya(true) jika != 0 maka bukan faktor nya(false)
		s = N%d == 0

		//output
		fmt.Println(d, s)
	}
}
