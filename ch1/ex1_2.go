package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, str := range os.Args[0:] {
		fmt.Printf("%d %s\n", idx, str)
	}
}
