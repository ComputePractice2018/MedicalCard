package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/medicalcard/backend/utils"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Привет, %s", "Мир")

	a := flag.Int("a", 1, "число 1")
	b := flag.Int("b", 2, "число 2")
	flag.Parse()

	var c = utils.Sum(*a, *b)
	fmt.Printf("\nСумма = %d", c)
}
