package main

import (
	"fmt"
	"time"
)

func panggilNama(nama string) {
	for i := 0; i < 5; i++ {
		//melakukan jeda sebanyak 100 milisecond
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hallo,", nama)
	}
}
func main() {
	fmt.Println("===== Program Goroutines (Async program) =====")
	go panggilNama("Bayu") // bagian ini bakalan dijalankan secara async. jadi bagian bawah bakalan dijalanin walaupun bagian ini belom selesai dijalanin
	panggilNama("Baki")

}
