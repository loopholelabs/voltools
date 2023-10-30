package fsdata

import (
	"errors"
	"io"
	"os"
	"sort"
)

var (
	ErrFilesizeNotSupported = errors.New("filesize not supported")
)

type fileData struct {
	size int64
	data []dataAt
}

type dataAt struct {
	loc  int64
	data []byte
}

func Format(gigs int, file io.WriterAt) error {
	data, ok := allFileData[gigs]
	if !ok {
		return ErrFilesizeNotSupported
	}

	for _, d := range data.data {
		if _, err := file.WriteAt(d.data, d.loc); err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(mbs int, file string) error {
	data, ok := allFileData[mbs]
	if !ok {
		return ErrFilesizeNotSupported
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	if err := f.Truncate(data.size); err != nil {
		return err
	}

	if err := Format(mbs, f); err != nil {
		return err
	}

	return f.Close()
}

func CreateNextLargestFile(mbs int, file string) error {
	keys := make([]int, 0, len(allFileData))
	for k := range allFileData {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Find the next largest size
	var nextLargestSize int
	found := false
	for _, k := range keys {
		if k >= mbs {
			nextLargestSize = k
			found = true
			break
		}
	}

	if !found {
		return ErrFilesizeNotSupported
	}

	return CreateFile(nextLargestSize, file)
}
