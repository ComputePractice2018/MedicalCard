package main

import "fmt"

// Sum суммирует два целых числа
func Sum(a, b int) int {
	var c = a + b
	return c
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Привет, %s", "Мир")

	var c = Sum(7, 8)
	fmt.Printf("\nСумма = %d", c)
}
