// package ini untuk menandakan bahwa ini adalah file main program dan akan di eksekusi pertama kali
package main

// import digunakan untuk memasukan modul lain kedalam file program
import "fmt"

// func yang akan dijalankan pertama kali saat file dijalankan
func main() {
	//contoh pembuatan variabel
	//menggunakan var
	var namaDepan string = "Bayu"
	var namaBelakang string
	namaBelakang = "Swari"

	//cara singkat deklarasi variabel
	nama := "Bayu Al ikhlas Swari" //tipe data akan menyesuaikan apa yang kita isi

	//deklarasi multi variabel
	var satu, dua, tiga = 1, 2, 3

	//jika ada variabel yang tidak dipakai bisa menggunakan _
	var _ = "ini ga bakalan di pake"

	//deklarasi variabel menggunakan new (pointer)
	contohPointer := new(string)
	fmt.Println(contohPointer)  // isinya alamat dari contohPointer
	fmt.Println(*contohPointer) //isinya data yang ditunjuk oleh pointer

	// deklarasi konstanta (nilai yang tidak bisa diubah)
	const phi float32 = 3.14
	fmt.Println(phi)

	//delarasi multi konstanta
	const (
		panjang = 10
		lebar   = 20
		bentuk  = "persegi"
	)

	fmt.Println(panjang, lebar, bentuk)

	// contoh penggunaan fmt
	fmt.Println("HALLO")
	fmt.Printf("Nama saya %s %s \n", namaDepan, namaBelakang) // printf untuk menampilkan variabel juga
	fmt.Printf("Hallo %s \n", nama)
	fmt.Printf("%d %d %d", satu, dua, tiga)
}
