package main

import "fmt"

func main() {
	fmt.Println("===== Program Function =====")
	cetak()
	cetakNama("Bayu")
	fmt.Println(tambah(100, 500))
	fmt.Println(tambahDanKali(10, 70))

}

// fungsi tanpa nilai balik dan tanpa parameter
func cetak() {
	fmt.Println("Hello Guys")
}

// fungsi tanpa nilai balik dan menggunakan parameter
func cetakNama(nama string) {
	fmt.Println("Hello", nama)
}

// fungsi dengan niilai balik dan menggunakan parameter
func tambah(a, b int32) int32 {
	return (a + b)
}

// fungsi mengembalikan multi nilai
func tambahDanKali(a, b int32) (int32, int32) {
	return a + b, a * b
}
