package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("===== Program Buffer =====")
	var buff bytes.Buffer

	// proses memasukan sebuah string ke buff
	fmt.Fprintf(&buff, "[Adegan %d] %s\n", 1, "Intro Video")

	//mencetak isi buff
	fmt.Println(buff.String())
	fmt.Println("Loading...")

	// menambahkan sebuah string ke buff
	fmt.Fprintf(&buff, "[Adegan %d], %s\n", 2, "Hallo Guyis")

	//mencetak isi buff
	fmt.Println(buff.String())

	buff.Write([]byte("kumar"))
	fmt.Printf("%b", buff.Len())
	fmt.Println(buff.String())
}
