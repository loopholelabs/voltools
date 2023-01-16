package main

import (
	"fmt"
	"os"
	"time"
)

type FileData struct {
	size int64
	data []DataAt
}

type DataAt struct {
	loc  int64
	data []byte
}

// Create a file from data
func CreateFile(data FileData, file string) error {
	ctime := time.Now()
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	err = f.Truncate(data.size)
	if err != nil {
		return err
	}

	writes := 0
	dataWritten := 0
	for _, d := range data.data {
		n, err := f.WriteAt(d.data, d.loc)
		if n != len(d.data) || err != nil {
			return err
		}
		writes++
		dataWritten += len(d.data)
	}
	f.Close()

	fmt.Printf("Filesystem %s size %d created in %d ms (%d writes, %d bytes)\n", file, data.size, time.Since(ctime).Milliseconds(), writes, dataWritten)
	return nil
}
