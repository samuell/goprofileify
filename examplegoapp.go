package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	fmt.Println("Writing 1M points:")
	for i := 0; i <= 1e6; i++ {
		fmt.Print(".")
	}
}
