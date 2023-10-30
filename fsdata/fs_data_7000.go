package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_7000_1024.bin
	locIndex7000LOC1024 []byte
	
	//go:embed fs_data_7000_37142528.bin
	locIndex7000LOC37142528 []byte
	
	//go:embed fs_data_7000_134217728.bin
	locIndex7000LOC134217728 []byte
	
	//go:embed fs_data_7000_402653184.bin
	locIndex7000LOC402653184 []byte
	
	//go:embed fs_data_7000_671088640.bin
	locIndex7000LOC671088640 []byte
	
	//go:embed fs_data_7000_939524096.bin
	locIndex7000LOC939524096 []byte
	
	//go:embed fs_data_7000_1207959552.bin
	locIndex7000LOC1207959552 []byte
	
	//go:embed fs_data_7000_2147483648.bin
	locIndex7000LOC2147483648 []byte
	
	//go:embed fs_data_7000_2281701376.bin
	locIndex7000LOC2281701376 []byte
	
	//go:embed fs_data_7000_3355443200.bin
	locIndex7000LOC3355443200 []byte
	
	//go:embed fs_data_7000_3623878656.bin
	locIndex7000LOC3623878656 []byte
	
	//go:embed fs_data_7000_4294967296.bin
	locIndex7000LOC4294967296 []byte
	
	//go:embed fs_data_7000_6442450944.bin
	locIndex7000LOC6442450944 []byte
	
	//go:embed fs_data_7000_6576668672.bin
	locIndex7000LOC6576668672 []byte
	
)

func init() {
 	allFileData[7000] = fileData{size:7340032000, data:[]dataAt{
		
			{loc:1024, data: locIndex7000LOC1024},
		
			{loc:37142528, data: locIndex7000LOC37142528},
		
			{loc:134217728, data: locIndex7000LOC134217728},
		
			{loc:402653184, data: locIndex7000LOC402653184},
		
			{loc:671088640, data: locIndex7000LOC671088640},
		
			{loc:939524096, data: locIndex7000LOC939524096},
		
			{loc:1207959552, data: locIndex7000LOC1207959552},
		
			{loc:2147483648, data: locIndex7000LOC2147483648},
		
			{loc:2281701376, data: locIndex7000LOC2281701376},
		
			{loc:3355443200, data: locIndex7000LOC3355443200},
		
			{loc:3623878656, data: locIndex7000LOC3623878656},
		
			{loc:4294967296, data: locIndex7000LOC4294967296},
		
			{loc:6442450944, data: locIndex7000LOC6442450944},
		
			{loc:6576668672, data: locIndex7000LOC6576668672},
		
	}}
}
