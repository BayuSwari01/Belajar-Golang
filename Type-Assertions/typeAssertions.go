package main

import "fmt"

func main() {
	fmt.Println("===== Program Type Assertions =====")

	var nama interface{} = "Bayu"

	n := nama.(string)

	fmt.Println(n)

	n, flag := nama.(string)

	fmt.Println(n, flag)

	na, flag := nama.(float64)

	fmt.Println(na, flag)

	nam := nama.(float64)

	fmt.Println(nam)
}
