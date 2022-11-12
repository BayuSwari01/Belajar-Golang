package main

import "fmt"

func main() {
	fmt.Println("===== Program Type Casting =====")
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Isi dari mean : %f \n", mean)
}
