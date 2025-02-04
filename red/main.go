package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.BgHiRed)
	out := ""
	for range 20 {
		for range 120 {
			out += red.Sprint(" ")
		}
		out += "\n"
	}
	fmt.Printf(out)
}
