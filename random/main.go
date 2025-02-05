package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	size := len(os.Args) - 1
	if size == 0 {
		fmt.Println("need at least one argument")
		os.Exit(1)
	}
	fmt.Println(os.Args[rand.IntN(size)+1])
}
