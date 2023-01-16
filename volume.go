package main

import (
	"fmt"
	"os"
)

type fileData struct {
	size int64
	data []dataAt
}

type dataAt struct {
	loc  int64
	data []byte
}

func CreateFile(gigs int, file string) error {
	data, ok := allFileData[gigs]
	if !ok {
		return fmt.Errorf("Filesize %dg is not supported.", gigs)
	}
	return createFile(data, file)
}

// Create a file from data
func createFile(data fileData, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	err = f.Truncate(data.size)
	if err != nil {
		return err
	}

	for _, d := range data.data {
		_, err := f.WriteAt(d.data, d.loc)
		if err != nil {
			return err
		}
	}

	return f.Close()
}
