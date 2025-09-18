package main

import "fmt"

func main() {
	var name string = "Dmytro"
	surname := "Luhinets"
	var age int = 35
	height := 183
	var writeInGo bool = true
	var result float64 = 45.34
	result2 := 456.776

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Surname: %s\n", surname)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Write in Go: %t\n", writeInGo)
	fmt.Printf("Result: %.2f\n", result)
	fmt.Printf("Result 2: %.2f\n", result2)
}