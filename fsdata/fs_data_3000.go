package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_3000_1025.bin
	locIndex3000LOC1025 []byte
	
	//go:embed fs_data_3000_34439168.bin
	locIndex3000LOC34439168 []byte
	
	//go:embed fs_data_3000_134217729.bin
	locIndex3000LOC134217729 []byte
	
	//go:embed fs_data_3000_402653185.bin
	locIndex3000LOC402653185 []byte
	
	//go:embed fs_data_3000_671088641.bin
	locIndex3000LOC671088641 []byte
	
	//go:embed fs_data_3000_939524097.bin
	locIndex3000LOC939524097 []byte
	
	//go:embed fs_data_3000_1207959553.bin
	locIndex3000LOC1207959553 []byte
	
	//go:embed fs_data_3000_1342177280.bin
	locIndex3000LOC1342177280 []byte
	
	//go:embed fs_data_3000_2147483648.bin
	locIndex3000LOC2147483648 []byte
	
)

func init() {
 	allFileData[3000] = fileData{size:3145728000, data:[]dataAt{
		
			{loc:1025, data: locIndex3000LOC1025},
		
			{loc:34439168, data: locIndex3000LOC34439168},
		
			{loc:134217729, data: locIndex3000LOC134217729},
		
			{loc:402653185, data: locIndex3000LOC402653185},
		
			{loc:671088641, data: locIndex3000LOC671088641},
		
			{loc:939524097, data: locIndex3000LOC939524097},
		
			{loc:1207959553, data: locIndex3000LOC1207959553},
		
			{loc:1342177280, data: locIndex3000LOC1342177280},
		
			{loc:2147483648, data: locIndex3000LOC2147483648},
		
	}}
}
