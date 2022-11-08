package main

import "fmt"

func main() {
	fmt.Println("===== Program percabangan Golang =====")

	var angka int16 = 10

	// if biasa
	if angka == 10 {
		fmt.Println("Lulus sempurna")
	} else if angka < 10 && angka >= 8 {
		fmt.Println("Lulus bagus")
	} else if angka < 8 && angka >= 6 {
		fmt.Println("Lulus pas pasan")
	} else {
		fmt.Println("Tidak lulus")
	}

	// variabel temporary di dalam if. variabel bonys hanya bisa digunakan didalam scope if statement
	if nilaiBonus := angka + 1; angka == 10 {
		fmt.Println("Lulus sempurna", nilaiBonus)
	} else if angka < 10 && angka >= 8 {
		fmt.Println("Lulus bagus", nilaiBonus)
	} else if angka < 8 && angka >= 6 {
		fmt.Println("Lulus pas pasan", nilaiBonus)
	} else {
		fmt.Println("Tidak lulus", nilaiBonus)
	}

	// switch case
	switch angka {
	case 10:
		fmt.Println("Sempurna")
	case 9:
		fmt.Println("Hampir sempurna")
	case 8, 7, 6:
		fmt.Println("Lulus")
	default:
		fmt.Println("Boleh la")
	}

	// switch case gaya if else
	switch {
	case angka > 10:
		fmt.Println("Lebih besar dari sepuluh")
	case angka > 5:
		fmt.Println("Lebih besar dari lima")
	}

	// switch case menggunakan fallthrough
	switch {
	case angka >= 10:
		fmt.Println("Lebih besar atau sama dengan sepuluh")
		fallthrough
	case angka > 5:
		fmt.Println("dan lebih besar dari lima")
	}
}
