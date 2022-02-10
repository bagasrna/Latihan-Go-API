package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/multiply"
)

func main() {
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(9, 9)

	hasilkali := multiply.Kali(9, 9)

	fmt.Println(result)

	fmt.Println(hasilkali)
}
