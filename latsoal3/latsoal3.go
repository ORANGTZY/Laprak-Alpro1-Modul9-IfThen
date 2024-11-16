package main

import "fmt"

func main() {
	var x, y int
	var isFactorXY, isFactorYX bool
	fmt.Print("Masukkan dua bilangan bulat positif (x y): ")
	fmt.Scanln(&x, &y)

	// Menentukan apakah x adalah faktor dari y
	if y%x == 0 {
		isFactorXY = true
	} else {
		isFactorXY = false
	}

	// Menentukan apakah y adalah faktor dari x
	if x%y == 0 {
		isFactorYX = true
	} else {
		isFactorYX = false
	}

	fmt.Println(isFactorXY)
	fmt.Println(isFactorYX)
}
