package main

import (
	"errors"
	"fmt"
	"strings"
)

// contoh penggunakan funsi sebagai error handling
func validate(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("tidak bisa kosong")
	} else {
		return nil
	}
}

// fungsi error handling yang memiliki 2 nilai kembalian
func cek(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak bisa kosong")
	} else {
		return true, nil
	}
}

func main() {
	fmt.Println("===== Program Error Handling =====")

	err := validate("")

	if err != nil {
		fmt.Println("Ada error : ", err)
	} else {
		fmt.Println("Tidak ada error")
	}

	hasil, err := cek("")
	if err != nil {
		fmt.Println("ada error", err, "tapi hasilnya", hasil)
	} else {
		fmt.Println("ga ada error", hasil)
	}

	// penggunaan panic. panic ini akan menghentikan program walaupun itu main program dan akan menampilkan persan error ke terminal
	if hasil, err := cek(""); hasil {
		fmt.Println("ga ada error", hasil)
	} else {
		panic(err.Error())
		fmt.Println("ada error", err, "tapi hasilnya", hasil)
	}
}
