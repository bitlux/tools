package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func main() {
	height := 20
	width := 120

	x, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("GetSize:", err)
	} else {
		width = x
	}

	if len(os.Args) == 2 {
		y, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%q is not a number: %v", os.Args[1], err)
		} else {
			height = y
		}
	}

	out := strings.Repeat(" ", width)
	out += strings.Repeat("\n", height-1)
	color.New(color.BgHiRed).Println(out)
}
