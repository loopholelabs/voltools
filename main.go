package main

import (
	"fmt"
	"github.com/loopholelabs/voltools/fsdata"
)

// Create a new volume

func main() {
	size := 10
	err := fsdata.CreateFile(size, fmt.Sprintf("volume.%dg", size))
	if err != nil {
		panic(err)
	}
}
