package main

import (
	"flag"

	"github.com/loopholelabs/voltools/fsdata"
)

// Create a new volume
func main() {
	size := flag.Int("size", 20000, "Size of the filesystem to create in MB")
	output := flag.String("output", "fs.img", "Path to write the filesystem to")

	flag.Parse()

	if err := fsdata.CreateNextLargestFile(*size, *output); err != nil {
		panic(err)
	}
}
