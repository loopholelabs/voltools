package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_2000_1025.bin
	locIndex2000LOC1025 []byte
	
	//go:embed fs_data_2000_33927168.bin
	locIndex2000LOC33927168 []byte
	
	//go:embed fs_data_2000_134217729.bin
	locIndex2000LOC134217729 []byte
	
	//go:embed fs_data_2000_402653185.bin
	locIndex2000LOC402653185 []byte
	
	//go:embed fs_data_2000_671088641.bin
	locIndex2000LOC671088641 []byte
	
	//go:embed fs_data_2000_805306368.bin
	locIndex2000LOC805306368 []byte
	
	//go:embed fs_data_2000_939524097.bin
	locIndex2000LOC939524097 []byte
	
	//go:embed fs_data_2000_1207959553.bin
	locIndex2000LOC1207959553 []byte
	
)

func init() {
 	allFileData[2000] = fileData{size:2097152000, data:[]dataAt{
		
			{loc:1025, data: locIndex2000LOC1025},
		
			{loc:33927168, data: locIndex2000LOC33927168},
		
			{loc:134217729, data: locIndex2000LOC134217729},
		
			{loc:402653185, data: locIndex2000LOC402653185},
		
			{loc:671088641, data: locIndex2000LOC671088641},
		
			{loc:805306368, data: locIndex2000LOC805306368},
		
			{loc:939524097, data: locIndex2000LOC939524097},
		
			{loc:1207959553, data: locIndex2000LOC1207959553},
		
	}}
}
