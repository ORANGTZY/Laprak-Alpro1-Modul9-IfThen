package main

import "fmt"

func main() {
	var a int
	var teks string
	fmt.Scan(&a)
	teks = "negative"
	if a >= 0 {
		teks = "positive"
	}
	fmt.Println(teks)
}
