package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	fmt.Println("Go PHP CC v0.0.1")
	if flag.NArg() > 0 {
		fmt.Printf("Hello, %s!\n", flag.Arg(0))
		return
	}

	fmt.Println("Hello, World!")
}
