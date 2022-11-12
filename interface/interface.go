package main

import "fmt"

type BangunDatar interface {
	Luas() float64
	Keliling() float64
}

type Persegi struct {
	Sisi float64
}

func (persegi Persegi) Luas() float64 {
	return persegi.Sisi * persegi.Sisi
}

func (persegi Persegi) Keliling() float64 {
	return persegi.Sisi * 4
}

type Lingkaran struct {
	Jarijari float64
}

func (lingkaran Lingkaran) Luas() float64 {
	return 3.14 * lingkaran.Jarijari * lingkaran.Jarijari
}

func (lingkaran Lingkaran) Keliling() float64 {
	return 2 * 3.14 * lingkaran.Jarijari
}

func measure(b BangunDatar) {
	fmt.Println(b)
	fmt.Println("Luas :", b.Luas())
	fmt.Println("Keliling :", b.Keliling())
}

func main() {
	fmt.Println("===== Program Interface =====")
	p := Persegi{Sisi: 10}
	l := Lingkaran{Jarijari: 70}

	measure(p)
	measure(l)
}
