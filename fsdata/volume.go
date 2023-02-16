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

type Formattable interface {
	io.WriterAt
	Truncate(size int64) error
}

func Format(gigs int, formattable Formattable) error {
	data, ok := allFileData[gigs]
	if !ok {
		return ErrFilesizeNotSupported
	}

	if err := formattable.Truncate(data.size); err != nil {
		return err
	}

	for _, d := range data.data {
		_, err := formattable.WriteAt(d.data, d.loc)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(gigs int, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	if err := Format(gigs, f); err != nil {
		return err
	}

	return f.Close()
}
