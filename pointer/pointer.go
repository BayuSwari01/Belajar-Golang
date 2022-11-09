package main

import "fmt"

func main() {
	fmt.Println("===== Program Pointer =====")
	var angka int = 8
	var pAngka *int = &angka

	fmt.Println(angka)
	fmt.Println(pAngka)  // nampilin alamat yang di tuju
	fmt.Println(*pAngka) //nampilin data dari alamat yang dituju

	*pAngka = 10 // mengganti variabel yg dituju oleh pointer menjadi 10
	fmt.Println(angka)

	ubah(pAngka)

	fmt.Println(angka)   //berubah menjadi 100
	fmt.Println(*pAngka) //berubah menjadi 100
}

// fungsi menggunakan pointer sebagai parameter. sehingga dapat merubah nilai tanpa membalikan nilai
func ubah(angka *int) {
	*angka = 100
}
