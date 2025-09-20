package main

import (
	"fmt"
	"os"

	"github.com/dluhinets/go-learning/internal/greeter"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeter.Hello(name))
	fmt.Println(greeter.Welcome("кросівки", 45))
}