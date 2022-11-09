package main

import "fmt"

func main() {
	fmt.Println("===== Program Array dan Perulangan =====")

	// perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke", i)
	}

	// perulangan for dengan argumen hanya kondisi
	j := 1
	for j < 5 {
		fmt.Println("Nilai j :", j)
		j++
	}

	// perulangan for tanpa argumen
	for {
		fmt.Println("Hallo")

		if j == 7 {
			j++
			continue //contoh penggunaan continue
		}

		if j > 10 {
			break
		}
		j++
	}

	// perulangan bersarang dan penggunaan outerloop
outerLoop:
	for i := 0; i < 10; i++ {
		for k := 0; k < 5; k++ {
			if k == 3 {
				break outerLoop
			}
			fmt.Println("ulang")
		}
	}

	//deklarasi array
	var arrNama [2]string
	arrNama[0] = "Bayu"
	arrNama[1] = "Baki"

	var arrNamaBelakang = [2]string{"Swari", "Super"}
	fmt.Println(arrNamaBelakang)

	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//deklarasi array multi dimensi
	var arrMulti = [2][3]int{{1, 2, 3}, {3, 2, 1}}

	fmt.Println(arrMulti)

	// perulangan berdasarkan isi array (kalo misal ada yg ga pengen dipake bisa pake variabel _ biar ga error)
	for i, nama := range arrNama {
		fmt.Println(i, nama)
	}

	//deklarasi array menggunakan make
	var buah = make([]string, 3)
	buah[0] = "mangga"
}
