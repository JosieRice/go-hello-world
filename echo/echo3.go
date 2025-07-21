package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // print all args
	fmt.Println(strings.Join(os.Args, " ")) // print all args and name of command
	fmt.Println(os.Args[1:]) // print all args in []

	// print all args on a new line with the index
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	} 
}