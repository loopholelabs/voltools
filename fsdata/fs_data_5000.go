package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_5000_1025.bin
	locIndex5000LOC1025 []byte
	
	//go:embed fs_data_5000_35463168.bin
	locIndex5000LOC35463168 []byte
	
	//go:embed fs_data_5000_134217729.bin
	locIndex5000LOC134217729 []byte
	
	//go:embed fs_data_5000_402653185.bin
	locIndex5000LOC402653185 []byte
	
	//go:embed fs_data_5000_671088641.bin
	locIndex5000LOC671088641 []byte
	
	//go:embed fs_data_5000_939524097.bin
	locIndex5000LOC939524097 []byte
	
	//go:embed fs_data_5000_1207959553.bin
	locIndex5000LOC1207959553 []byte
	
	//go:embed fs_data_5000_2147483648.bin
	locIndex5000LOC2147483648 []byte
	
	//go:embed fs_data_5000_2281701376.bin
	locIndex5000LOC2281701376 []byte
	
	//go:embed fs_data_5000_3355443201.bin
	locIndex5000LOC3355443201 []byte
	
	//go:embed fs_data_5000_3623878657.bin
	locIndex5000LOC3623878657 []byte
	
	//go:embed fs_data_5000_4294967296.bin
	locIndex5000LOC4294967296 []byte
	
)

func init() {
 	allFileData[5000] = fileData{size:5242880000, data:[]dataAt{
		
			{loc:1025, data: locIndex5000LOC1025},
		
			{loc:35463168, data: locIndex5000LOC35463168},
		
			{loc:134217729, data: locIndex5000LOC134217729},
		
			{loc:402653185, data: locIndex5000LOC402653185},
		
			{loc:671088641, data: locIndex5000LOC671088641},
		
			{loc:939524097, data: locIndex5000LOC939524097},
		
			{loc:1207959553, data: locIndex5000LOC1207959553},
		
			{loc:2147483648, data: locIndex5000LOC2147483648},
		
			{loc:2281701376, data: locIndex5000LOC2281701376},
		
			{loc:3355443201, data: locIndex5000LOC3355443201},
		
			{loc:3623878657, data: locIndex5000LOC3623878657},
		
			{loc:4294967296, data: locIndex5000LOC4294967296},
		
	}}
}
