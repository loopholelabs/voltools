package fsdata

import _ "embed"

{{ $index := .Index }}

var (
	{{ range $loc := .Chunks }}
	//go:embed fs_data_{{ $index }}_{{ $loc }}.bin
	locIndex{{ $index }}LOC{{ $loc }} []byte
	{{ end }}
)

func init() {
 	allFileData[{{ $index }}] = fileData{size:{{ .Size }}, data:[]dataAt{
		{{ range $loc := .Chunks }}
			{loc:{{ $loc }}, data: locIndex{{ $index }}LOC{{ $loc }}},
		{{ end }}
	}}
}
