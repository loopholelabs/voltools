package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

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

var prefix string

func flush() int {
	// Output current data

	// Strip trailing 0s
	l := len(data)
	for {
		if l == 0 {
			return -1 // No data
		}
		if data[l-1] != 0 {
			break
		}
		l--
	}
	data = data[0:l]

	dataFile, err := os.OpenFile(fmt.Sprintf("%v%v.bin", prefix, dataloc), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	if _, err := dataFile.Write(data); err != nil {
		panic(err)
	}
	data = make([]byte, 0)

	rv := dataloc

	dataloc = 0

	return rv
}

var (
	//go:embed fs_data.go.tpl
	fsDataTemplate string

	//go:embed fs_data_00_init.go.tpl
	fsDataInitTemplate []byte

	//go:embed volume.go.tpl
	volumeTemplate []byte
)

type templateData struct {
	Chunks []int
	Index  int
	Size   int64
}

func main() {
	input := flag.String("input", "", "Input file path")
	size := flag.Int("size", 0, "Input file size in GB")
	outputDir := flag.String("output-dir", "out", "Output directory")
	maxZeros := flag.Int("max-zeros", 1024*1024, "Maximum count of blocks of zeros to include before terminating a section")

	flag.Parse()

	log.Printf("Processing %v (%v GB)", *input, *size)

	data = make([]byte, 0)
	dataloc = 0
	prefix = filepath.Join(*outputDir, fmt.Sprintf("fs_data_%v_", *size))

	inputFile, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	stat, err := inputFile.Stat()
	if err != nil {
		panic(err)
	}

	location := 0

	buffer := make([]byte, 1024*1024)

	numzeros := 0

	d := templateData{
		Chunks: []int{},
		Index:  *size,
		Size:   stat.Size(),
	}

	for {
		n, err := inputFile.Read(buffer)
		for i := 0; i < n; i++ {
			b := buffer[i]
			if b == 0 {
				numzeros++
				addByte(location, b)

				if numzeros > *maxZeros {
					loc := flush()
					if loc > -1 {
						d.Chunks = append(d.Chunks, loc)
					}
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

	loc := flush()
	if loc > -1 {
		d.Chunks = append(d.Chunks, loc)
	}

	{
		outputFile, err := os.OpenFile(filepath.Join(*outputDir, fmt.Sprintf("fs_data_%v.go", *size)), os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		tpl, err := template.New("fs_data.go.tpl").Parse(fsDataTemplate)
		if err != nil {
			panic(err)
		}

		if err := tpl.Execute(outputFile, d); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(filepath.Join(*outputDir, "fs_data_00_init.go"), fsDataInitTemplate, 0644); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(*outputDir, "volume.go"), volumeTemplate, 0644); err != nil {
		panic(err)
	}
}
