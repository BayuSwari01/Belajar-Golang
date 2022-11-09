package main

import "fmt"

//mendefinisikan sebuah struct

type Orang struct {
	namaDepan    string
	namaBelakang string
}

func main() {
	fmt.Println("===== Program Struct =====")

	// membuat variabel dengan tipe data struct yang sudah dibuat
	bayu := Orang{namaDepan: "Bayu", namaBelakang: "Swari"}
	var baki Orang
	baki.namaDepan = "Baki"
	baki.namaBelakang = "Super"

	fmt.Println(bayu, baki)

	//mengubah data
	baki.namaBelakang = "Swari"
	fmt.Println(bayu, baki)
}
