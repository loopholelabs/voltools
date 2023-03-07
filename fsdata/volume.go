package fsdata

import (
	"errors"
	"io"
	"os"
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

func CreateFile(gigs int, file string) error {
	data, ok := allFileData[gigs]
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

	if err := Format(gigs, f); err != nil {
		return err
	}

	return f.Close()
}
