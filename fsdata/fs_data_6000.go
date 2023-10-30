package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_6000_1024.bin
	locIndex6000LOC1024 []byte
	
	//go:embed fs_data_6000_36696064.bin
	locIndex6000LOC36696064 []byte
	
	//go:embed fs_data_6000_134217728.bin
	locIndex6000LOC134217728 []byte
	
	//go:embed fs_data_6000_402653184.bin
	locIndex6000LOC402653184 []byte
	
	//go:embed fs_data_6000_671088640.bin
	locIndex6000LOC671088640 []byte
	
	//go:embed fs_data_6000_939524096.bin
	locIndex6000LOC939524096 []byte
	
	//go:embed fs_data_6000_1207959552.bin
	locIndex6000LOC1207959552 []byte
	
	//go:embed fs_data_6000_2147483648.bin
	locIndex6000LOC2147483648 []byte
	
	//go:embed fs_data_6000_2281701376.bin
	locIndex6000LOC2281701376 []byte
	
	//go:embed fs_data_6000_3355443200.bin
	locIndex6000LOC3355443200 []byte
	
	//go:embed fs_data_6000_3623878656.bin
	locIndex6000LOC3623878656 []byte
	
	//go:embed fs_data_6000_4294967296.bin
	locIndex6000LOC4294967296 []byte
	
)

func init() {
 	allFileData[6000] = fileData{size:6291456000, data:[]dataAt{
		
			{loc:1024, data: locIndex6000LOC1024},
		
			{loc:36696064, data: locIndex6000LOC36696064},
		
			{loc:134217728, data: locIndex6000LOC134217728},
		
			{loc:402653184, data: locIndex6000LOC402653184},
		
			{loc:671088640, data: locIndex6000LOC671088640},
		
			{loc:939524096, data: locIndex6000LOC939524096},
		
			{loc:1207959552, data: locIndex6000LOC1207959552},
		
			{loc:2147483648, data: locIndex6000LOC2147483648},
		
			{loc:2281701376, data: locIndex6000LOC2281701376},
		
			{loc:3355443200, data: locIndex6000LOC3355443200},
		
			{loc:3623878656, data: locIndex6000LOC3623878656},
		
			{loc:4294967296, data: locIndex6000LOC4294967296},
		
	}}
}
