package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_5500_1025.bin
	locIndex5500LOC1025 []byte
	
	//go:embed fs_data_5500_36507648.bin
	locIndex5500LOC36507648 []byte
	
	//go:embed fs_data_5500_134217729.bin
	locIndex5500LOC134217729 []byte
	
	//go:embed fs_data_5500_402653185.bin
	locIndex5500LOC402653185 []byte
	
	//go:embed fs_data_5500_671088641.bin
	locIndex5500LOC671088641 []byte
	
	//go:embed fs_data_5500_939524097.bin
	locIndex5500LOC939524097 []byte
	
	//go:embed fs_data_5500_1207959553.bin
	locIndex5500LOC1207959553 []byte
	
	//go:embed fs_data_5500_2147483648.bin
	locIndex5500LOC2147483648 []byte
	
	//go:embed fs_data_5500_2281701376.bin
	locIndex5500LOC2281701376 []byte
	
	//go:embed fs_data_5500_3355443201.bin
	locIndex5500LOC3355443201 []byte
	
	//go:embed fs_data_5500_3623878657.bin
	locIndex5500LOC3623878657 []byte
	
	//go:embed fs_data_5500_4294967296.bin
	locIndex5500LOC4294967296 []byte
	
)

func init() {
 	allFileData[5500] = fileData{size:5767168000, data:[]dataAt{
		
			{loc:1025, data: locIndex5500LOC1025},
		
			{loc:36507648, data: locIndex5500LOC36507648},
		
			{loc:134217729, data: locIndex5500LOC134217729},
		
			{loc:402653185, data: locIndex5500LOC402653185},
		
			{loc:671088641, data: locIndex5500LOC671088641},
		
			{loc:939524097, data: locIndex5500LOC939524097},
		
			{loc:1207959553, data: locIndex5500LOC1207959553},
		
			{loc:2147483648, data: locIndex5500LOC2147483648},
		
			{loc:2281701376, data: locIndex5500LOC2281701376},
		
			{loc:3355443201, data: locIndex5500LOC3355443201},
		
			{loc:3623878657, data: locIndex5500LOC3623878657},
		
			{loc:4294967296, data: locIndex5500LOC4294967296},
		
	}}
}
