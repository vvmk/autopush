package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		watch(os.Args[1])
	} else {
		watch("")
	}
}
