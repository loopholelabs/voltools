package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_8000_1024.bin
	locIndex8000LOC1024 []byte
	
	//go:embed fs_data_8000_37523456.bin
	locIndex8000LOC37523456 []byte
	
	//go:embed fs_data_8000_134217728.bin
	locIndex8000LOC134217728 []byte
	
	//go:embed fs_data_8000_402653184.bin
	locIndex8000LOC402653184 []byte
	
	//go:embed fs_data_8000_671088640.bin
	locIndex8000LOC671088640 []byte
	
	//go:embed fs_data_8000_939524096.bin
	locIndex8000LOC939524096 []byte
	
	//go:embed fs_data_8000_1207959552.bin
	locIndex8000LOC1207959552 []byte
	
	//go:embed fs_data_8000_2147483648.bin
	locIndex8000LOC2147483648 []byte
	
	//go:embed fs_data_8000_2281701376.bin
	locIndex8000LOC2281701376 []byte
	
	//go:embed fs_data_8000_3355443200.bin
	locIndex8000LOC3355443200 []byte
	
	//go:embed fs_data_8000_3623878656.bin
	locIndex8000LOC3623878656 []byte
	
	//go:embed fs_data_8000_4294967296.bin
	locIndex8000LOC4294967296 []byte
	
	//go:embed fs_data_8000_6442450944.bin
	locIndex8000LOC6442450944 []byte
	
	//go:embed fs_data_8000_6576668672.bin
	locIndex8000LOC6576668672 []byte
	
)

func init() {
 	allFileData[8000] = fileData{size:8388608000, data:[]dataAt{
		
			{loc:1024, data: locIndex8000LOC1024},
		
			{loc:37523456, data: locIndex8000LOC37523456},
		
			{loc:134217728, data: locIndex8000LOC134217728},
		
			{loc:402653184, data: locIndex8000LOC402653184},
		
			{loc:671088640, data: locIndex8000LOC671088640},
		
			{loc:939524096, data: locIndex8000LOC939524096},
		
			{loc:1207959552, data: locIndex8000LOC1207959552},
		
			{loc:2147483648, data: locIndex8000LOC2147483648},
		
			{loc:2281701376, data: locIndex8000LOC2281701376},
		
			{loc:3355443200, data: locIndex8000LOC3355443200},
		
			{loc:3623878656, data: locIndex8000LOC3623878656},
		
			{loc:4294967296, data: locIndex8000LOC4294967296},
		
			{loc:6442450944, data: locIndex8000LOC6442450944},
		
			{loc:6576668672, data: locIndex8000LOC6576668672},
		
	}}
}
