package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_2500_1025.bin
	locIndex2500LOC1025 []byte
	
	//go:embed fs_data_2500_34185216.bin
	locIndex2500LOC34185216 []byte
	
	//go:embed fs_data_2500_134217729.bin
	locIndex2500LOC134217729 []byte
	
	//go:embed fs_data_2500_402653185.bin
	locIndex2500LOC402653185 []byte
	
	//go:embed fs_data_2500_671088641.bin
	locIndex2500LOC671088641 []byte
	
	//go:embed fs_data_2500_939524097.bin
	locIndex2500LOC939524097 []byte
	
	//go:embed fs_data_2500_1073741824.bin
	locIndex2500LOC1073741824 []byte
	
	//go:embed fs_data_2500_1207959553.bin
	locIndex2500LOC1207959553 []byte
	
	//go:embed fs_data_2500_2147483648.bin
	locIndex2500LOC2147483648 []byte
	
)

func init() {
 	allFileData[2500] = fileData{size:2621440000, data:[]dataAt{
		
			{loc:1025, data: locIndex2500LOC1025},
		
			{loc:34185216, data: locIndex2500LOC34185216},
		
			{loc:134217729, data: locIndex2500LOC134217729},
		
			{loc:402653185, data: locIndex2500LOC402653185},
		
			{loc:671088641, data: locIndex2500LOC671088641},
		
			{loc:939524097, data: locIndex2500LOC939524097},
		
			{loc:1073741824, data: locIndex2500LOC1073741824},
		
			{loc:1207959553, data: locIndex2500LOC1207959553},
		
			{loc:2147483648, data: locIndex2500LOC2147483648},
		
	}}
}
