package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan bulat positif (x y): ")
	fmt.Scanln(&x, &y)

	// Menentukan apakah x adalah faktor dari y
	var isFactorXY bool
	if y%x == 0 {
		isFactorXY = true
	} else {
		isFactorXY = false
	}

	// Menentukan apakah y adalah faktor dari x
	var isFactorYX bool
	if x%y == 0 {
		isFactorYX = true
	} else {
		isFactorYX = false
	}

	fmt.Println(isFactorXY)
	fmt.Println(isFactorYX)
}
