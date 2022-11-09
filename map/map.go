package main

import "fmt"

func main() {
	fmt.Println("===== Program Map =====")

	//deklarasi map
	var mapNama = map[string]string{}

	mapNama["depan"] = "Bayu"
	mapNama["belakang"] = "Swari"

	fmt.Println(mapNama)

	// perulangan untuk mengakses map
	for key, val := range mapNama {
		fmt.Println(key, val)
	}

	// menghapus item map
	delete(mapNama, "belakang")

	fmt.Println(mapNama)

	// mengecek apakah key ada didalam map
	if value, isExist := mapNama["belakang"]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Tidak ada")
	}

	//penggabungan antara slice dan map
	var slmapNama = []map[string]string{
		{"depan": "Bayu", "belakang": "Swari"},
		{"depan": "Baki", "belakang": "Swari"},
	}

	fmt.Println(slmapNama)
}
