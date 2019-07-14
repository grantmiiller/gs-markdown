package main

import (
	"fmt"
	"os"

	"github.com/grantmiiller/gs-mkdown/token"
)

func main() {
	t, err := token.New("EOF", "WHOOP")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	t.ToString()
	fmt.Println("Hello, world.")
}
