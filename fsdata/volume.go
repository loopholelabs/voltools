package fsdata

import (
	"errors"
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

type formattableFile struct {
	*os.File
}

func (f formattableFile) WriteAt(chunk []byte, offset int64) (err error) {
	_, err = f.File.WriteAt(chunk, offset)

	return
}

type Formattable interface {
	WriteAt(chunk []byte, offset int64) error
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
		if err := formattable.WriteAt(d.data, d.loc); err != nil {
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

	if err := Format(gigs, &formattableFile{f}); err != nil {
		return err
	}

	return f.Close()
}
