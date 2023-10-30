package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_4500_1025.bin
	locIndex4500LOC1025 []byte
	
	//go:embed fs_data_4500_35209216.bin
	locIndex4500LOC35209216 []byte
	
	//go:embed fs_data_4500_134217729.bin
	locIndex4500LOC134217729 []byte
	
	//go:embed fs_data_4500_402653185.bin
	locIndex4500LOC402653185 []byte
	
	//go:embed fs_data_4500_671088641.bin
	locIndex4500LOC671088641 []byte
	
	//go:embed fs_data_4500_939524097.bin
	locIndex4500LOC939524097 []byte
	
	//go:embed fs_data_4500_1207959553.bin
	locIndex4500LOC1207959553 []byte
	
	//go:embed fs_data_4500_2147483648.bin
	locIndex4500LOC2147483648 []byte
	
	//go:embed fs_data_4500_2281701376.bin
	locIndex4500LOC2281701376 []byte
	
	//go:embed fs_data_4500_3355443201.bin
	locIndex4500LOC3355443201 []byte
	
	//go:embed fs_data_4500_3623878657.bin
	locIndex4500LOC3623878657 []byte
	
	//go:embed fs_data_4500_4294967296.bin
	locIndex4500LOC4294967296 []byte
	
)

func init() {
 	allFileData[4500] = fileData{size:4718592000, data:[]dataAt{
		
			{loc:1025, data: locIndex4500LOC1025},
		
			{loc:35209216, data: locIndex4500LOC35209216},
		
			{loc:134217729, data: locIndex4500LOC134217729},
		
			{loc:402653185, data: locIndex4500LOC402653185},
		
			{loc:671088641, data: locIndex4500LOC671088641},
		
			{loc:939524097, data: locIndex4500LOC939524097},
		
			{loc:1207959553, data: locIndex4500LOC1207959553},
		
			{loc:2147483648, data: locIndex4500LOC2147483648},
		
			{loc:2281701376, data: locIndex4500LOC2281701376},
		
			{loc:3355443201, data: locIndex4500LOC3355443201},
		
			{loc:3623878657, data: locIndex4500LOC3623878657},
		
			{loc:4294967296, data: locIndex4500LOC4294967296},
		
	}}
}
