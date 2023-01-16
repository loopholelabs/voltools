package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const MAX_ZEROS = 32

var data []byte
var dataloc int

func addByte(location int, newdata byte) {
	if len(data) == 0 {
		if newdata == 0 {
			return
		}
		dataloc = location
	}
	data = append(data, newdata)
}

var is_first = true

func flush() {
	// Output current data

	// Strip trailing 0s
	l := len(data)
	for {
		if l == 0 {
			return // No data
		}
		if data[l-1] != 0 {
			break
		}
		l--
	}
	data = data[0:l]

	if !is_first {
		fmt.Printf(", ")
	}
	is_first = false
	fmt.Printf("{loc:%d, data:[]byte{", dataloc)
	for i, b := range data {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", b)
	}

	fmt.Printf("}}")
	//, hex.EncodeToString(data))
	data = make([]byte, 0)
	dataloc = 0
}

func main() {
	fname := os.Args[1]
	gigs, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n// Processing %s which is %d gigs...\n", fname, gigs)

	data = make([]byte, 0)
	dataloc = 0

	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf(" allFileData[%d] = fileData{size:%d, data:[]dataAt{", gigs, stat.Size())

	location := 0

	buffer := make([]byte, 1024*1024)

	numzeros := 0

	for {
		n, err := f.Read(buffer)
		for i := 0; i < n; i++ {
			b := buffer[i]
			if b == 0 {
				numzeros++
				addByte(location, b)

				if numzeros > MAX_ZEROS {
					flush()
				}
			} else {
				numzeros = 0
				addByte(location, b)
			}
			location++
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	flush()
	fmt.Printf("}}\n")
}
