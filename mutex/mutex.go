package main

import (
	"fmt"
	"sync"
)

type counter struct {
	//menambahkan objek mutex
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.Lock() //mengunci variabel agar tidak diakses secara bersamaan
	c.val++
	c.Unlock() //membuka kunci agar variabel dapat diakses oleh fungsi yang lain
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	fmt.Println("===== Program Mutex =====")
	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
