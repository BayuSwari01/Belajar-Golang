package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("===== Program Buffer =====")

	var strByyte = []byte("Hello, Ranjan, Kumar Good we will meet again,")

	// memotong string bagian depan jika ada kata "Hello, "
	strByyte = bytes.TrimPrefix(strByyte, []byte("Hello, "))

	// memotong string bagian belakang jika ada kata "Good we will meet again,"
	strByyte = bytes.TrimSuffix(strByyte, []byte("Good we will meet again,"))
	fmt.Printf("Hello %s \n", strByyte)
}
