package main

import "fmt"

func main() {
	var name string = "Dmytro"
	surname := "Luhinets"
	var age int = 35
	stature := 183
	var writeInGo bool = true
	var result float64 = 45.34
	result2 := 456.776

	fmt.Println("Name: " + name)
	fmt.Println("Surname: " + surname)
	fmt.Println("Surname: ", age)
	fmt.Println("Stature: ", stature)
	fmt.Println("Write in Go: ", writeInGo)
	fmt.Println("Result: ", result)
	fmt.Println("Result 2: ", result2)
}