package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_1000_1025.bin
	locIndex1000LOC1025 []byte
	
	//go:embed fs_data_1000_16965632.bin
	locIndex1000LOC16965632 []byte
	
	//go:embed fs_data_1000_134217729.bin
	locIndex1000LOC134217729 []byte
	
	//go:embed fs_data_1000_268435456.bin
	locIndex1000LOC268435456 []byte
	
	//go:embed fs_data_1000_402653185.bin
	locIndex1000LOC402653185 []byte
	
	//go:embed fs_data_1000_671088641.bin
	locIndex1000LOC671088641 []byte
	
	//go:embed fs_data_1000_939524097.bin
	locIndex1000LOC939524097 []byte
	
)

func init() {
 	allFileData[1000] = fileData{size:1048576000, data:[]dataAt{
		
			{loc:1025, data: locIndex1000LOC1025},
		
			{loc:16965632, data: locIndex1000LOC16965632},
		
			{loc:134217729, data: locIndex1000LOC134217729},
		
			{loc:268435456, data: locIndex1000LOC268435456},
		
			{loc:402653185, data: locIndex1000LOC402653185},
		
			{loc:671088641, data: locIndex1000LOC671088641},
		
			{loc:939524097, data: locIndex1000LOC939524097},
		
	}}
}
