package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak bisa kosong")
	}
	return true, nil
}

func jalan() {
	fmt.Println("===== Program Error Handling =====")

	defer catch()

	if hasil, err := validate(""); hasil {
		fmt.Println("ga ada error", hasil)
	} else {
		panic(err.Error())
		fmt.Println("ada error", err, "tapi hasilnya", hasil)
	}
}

func main() {
	jalan()
	fmt.Println("masih jalan") //ini masih jalan karena panic hanya menghentikan fungsi jalan saja. ketika fungsi jalan selesai maka akan tetap lanjut kebawah lagi
}
