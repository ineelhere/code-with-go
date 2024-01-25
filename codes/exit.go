package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	// The following line won't be executed due to os.Exit.
	os.Exit(3)
}
