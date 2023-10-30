package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_4000_1025.bin
	locIndex4000LOC1025 []byte
	
	//go:embed fs_data_4000_34951168.bin
	locIndex4000LOC34951168 []byte
	
	//go:embed fs_data_4000_134217729.bin
	locIndex4000LOC134217729 []byte
	
	//go:embed fs_data_4000_402653185.bin
	locIndex4000LOC402653185 []byte
	
	//go:embed fs_data_4000_671088641.bin
	locIndex4000LOC671088641 []byte
	
	//go:embed fs_data_4000_939524097.bin
	locIndex4000LOC939524097 []byte
	
	//go:embed fs_data_4000_1207959553.bin
	locIndex4000LOC1207959553 []byte
	
	//go:embed fs_data_4000_1879048192.bin
	locIndex4000LOC1879048192 []byte
	
	//go:embed fs_data_4000_2147483648.bin
	locIndex4000LOC2147483648 []byte
	
	//go:embed fs_data_4000_3355443201.bin
	locIndex4000LOC3355443201 []byte
	
	//go:embed fs_data_4000_3623878657.bin
	locIndex4000LOC3623878657 []byte
	
)

func init() {
 	allFileData[4000] = fileData{size:4194304000, data:[]dataAt{
		
			{loc:1025, data: locIndex4000LOC1025},
		
			{loc:34951168, data: locIndex4000LOC34951168},
		
			{loc:134217729, data: locIndex4000LOC134217729},
		
			{loc:402653185, data: locIndex4000LOC402653185},
		
			{loc:671088641, data: locIndex4000LOC671088641},
		
			{loc:939524097, data: locIndex4000LOC939524097},
		
			{loc:1207959553, data: locIndex4000LOC1207959553},
		
			{loc:1879048192, data: locIndex4000LOC1879048192},
		
			{loc:2147483648, data: locIndex4000LOC2147483648},
		
			{loc:3355443201, data: locIndex4000LOC3355443201},
		
			{loc:3623878657, data: locIndex4000LOC3623878657},
		
	}}
}
