package main

import (
	"fmt"
)

// Create a new volume

func main() {
	size := 1
	err := CreateFile(size, fmt.Sprintf("volume.%dg", size))
	if err != nil {
		panic(err)
	}
}
