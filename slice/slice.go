package main

import "fmt"

func main() {
	fmt.Println("===== Program Slice =====")

	//pembuatan slice mirip dengan array, tetapi bedanya disini tidak memerlukan ukuran dari slice
	var slNama = []string{"Bayu", "Baki", "Budi"}

	fmt.Println(slNama)

	//teknik pengaksesan 2 indeks yang hanya bisa dilakukan slice
	fmt.Println(slNama[1:3])

	// membuat slice dari potongan slice
	var potNama = slNama[0:2]

	fmt.Println(potNama)

	//slice merupakan tipe data reference (kaya pointer) jadi ketika kita ubah maka referensinya akan berubah juga
	potNama[0] = "Swari"

	fmt.Println(slNama) // hasilnya "Swari, Baki, Budi" karena potNama mereferensi dari slNama ketika potNama dirumah maka slNama juga akan berpengaruh

	// fungsi dari slice

	//len() untuk mengetahui panjang dari slice
	fmt.Println(len(slNama))

	//cap() untuk mengetahui kapasistas maksimum dari slice
	fmt.Println(cap(slNama))

	//append() untuk menambah element pada slice
	var baru = append(slNama, "Asep") // baru merupakan slice, jadi dia mereferensi juga ke slice slNama

	fmt.Println(baru)

	//copy() untuk mengcopy slice. untuk ukuran tetap mengikuti ukuran slice tujuan
	contoh := make([]string, 2)
	n := copy(contoh, slNama)

	fmt.Println(n)      // hasilnya 2
	fmt.Println(contoh) // hasilnya indeks ke 0 dan 1 dari slNama

	// pengaksesan menggunakan 3 indeks
	var tigaIndeks = slNama[0:2:3] // angka terakhir untuk ukuran dari si slice baru

	fmt.Println(tigaIndeks)
	fmt.Println(len(tigaIndeks))
	fmt.Println(cap(tigaIndeks))
}
