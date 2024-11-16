package main

import "fmt"

func main() {
	var a int
	var hasil bool
	fmt.Scan(&a)
	// if a < 0 && a%2 == 0 {
	// 	hasil = true
	// }
	hasil = a < 0 && a%2 == 0
	fmt.Print(hasil)
}
