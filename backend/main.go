package main

import (
	"fmt"

	"github.com/ComputePractice2018/medicalcard/backend/utils"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Привет, %s", "Мир")

	var c = utils.Sum(9, 8)
	fmt.Printf("\nСумма = %d", c)
}
