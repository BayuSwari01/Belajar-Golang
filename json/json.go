package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
	Kota string `json:"kota"`
}

func main() {
	fmt.Println("===== Program Membaca dan Menulis JSON =====")
	//membuat data json
	dataJson := `{"nama":"Bayu", "umur":17, "kota":"Bogor"}`

	//mengkonvert json menjadi byte
	dataJsonByte := []byte(dataJson)

	var user User

	// mengisikan data json tadi ke struct user (mengubah json menjadi struct)
	err := json.Unmarshal(dataJsonByte, &user)

	//jika terjadi error maka akan masuk ke sini
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//mencetak data
	fmt.Println("Data asli : ", dataJson)
	fmt.Println("Data hasil decode : ", user)
	fmt.Println("untuk mencetak data secara spesifik, misal mau cetak nama maka tinggal kasih .(nama variabelnya)", user.Nama)

	//mengubah stuct menjadi json
	keJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//mengubah json tadi ke string agar dapat dicetak / dicek
	balikKeStringUntukDicetak := string(keJson)

	//hasilnya
	fmt.Println("hasil :", balikKeStringUntukDicetak)
}
