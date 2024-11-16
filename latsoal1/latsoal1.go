package main

import "fmt"

func main() {
    var jumlahOrang, jumlahMotor int
    fmt.Print("Masukkan jumlah orang yang akan touring: ")
    fmt.Scan(&jumlahOrang)

    if jumlahOrang%2 == 0 {
        jumlahMotor = jumlahOrang / 2
    } else {
        jumlahMotor = (jumlahOrang / 2) + 1
    }

    fmt.Println("Jumlah motor yang diperlukan:", jumlahMotor)
}
