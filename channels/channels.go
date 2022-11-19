package main

import (
	"fmt"
	"time"
)

func cegah(nama string, c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(nama + " jangan mencuri")
	}

	// setelah melewati dibawah ini maka chanel akan menganggap bahwa fungsi ini telah selesai dijalankan dan di main fungsi akan mencetak hasilnya
	c <- nama + " tidak jadi mencuri"
}

func main() {
	fmt.Println("===== Program Channels =====")
	c := make(chan string)
	d := make(chan string)
	// kedua fungsi dibawah akan dijalankan secara async
	go cegah("swiper", c)
	go cegah("baki", d)

	//tetapi baris dibawah ini tidak akan berjalan sebelum fungsi async diatas selesai. hal ini dikarenakan ada variabel chanel
	fmt.Println(<-c)
	fmt.Println(<-d)
}
