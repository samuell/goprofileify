package main

import (
	"fmt"
)

func main() {
	fmt.Println("Writing 1M points:")
	for i := 0; i <= 1e6; i++ {
		fmt.Print(".")
	}
}
