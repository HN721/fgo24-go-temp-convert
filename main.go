package main

import (
	"fmt"
	"temp/calculate"
	"temp/history"
)

func main() {
	var choice int
	for {
		fmt.Println("====Welcome to converter Suhu====")
		fmt.Println("1. Konversi ")
		fmt.Println("2. History")
		fmt.Println("Pilih : 1-2")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var c float64
			fmt.Print("Masukkan suhu dalam Celsius: ")
			fmt.Scan(&c)
			k := calculate.CtoK(c)
			fmt.Printf("%.2f°C = %.2fK\n", c, k)

			history.AddHistory(&c, &k)
			fmt.Println("Tekan 1 untuk Konversi lagi")
			fmt.Scan(&choice)

		case 2:
			history.History()
			fmt.Println("Tekan 1 untuk konversi lagi")
			fmt.Scan(&choice)

		default:
			fmt.Println("Pilihan tidak valid!")
		}

	}
}
