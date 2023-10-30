package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_7500_1024.bin
	locIndex7500LOC1024 []byte
	
	//go:embed fs_data_7500_37335040.bin
	locIndex7500LOC37335040 []byte
	
	//go:embed fs_data_7500_134217728.bin
	locIndex7500LOC134217728 []byte
	
	//go:embed fs_data_7500_402653184.bin
	locIndex7500LOC402653184 []byte
	
	//go:embed fs_data_7500_671088640.bin
	locIndex7500LOC671088640 []byte
	
	//go:embed fs_data_7500_939524096.bin
	locIndex7500LOC939524096 []byte
	
	//go:embed fs_data_7500_1207959552.bin
	locIndex7500LOC1207959552 []byte
	
	//go:embed fs_data_7500_2147483648.bin
	locIndex7500LOC2147483648 []byte
	
	//go:embed fs_data_7500_2281701376.bin
	locIndex7500LOC2281701376 []byte
	
	//go:embed fs_data_7500_3355443200.bin
	locIndex7500LOC3355443200 []byte
	
	//go:embed fs_data_7500_3623878656.bin
	locIndex7500LOC3623878656 []byte
	
	//go:embed fs_data_7500_4294967296.bin
	locIndex7500LOC4294967296 []byte
	
	//go:embed fs_data_7500_6442450944.bin
	locIndex7500LOC6442450944 []byte
	
	//go:embed fs_data_7500_6576668672.bin
	locIndex7500LOC6576668672 []byte
	
)

func init() {
 	allFileData[7500] = fileData{size:7864320000, data:[]dataAt{
		
			{loc:1024, data: locIndex7500LOC1024},
		
			{loc:37335040, data: locIndex7500LOC37335040},
		
			{loc:134217728, data: locIndex7500LOC134217728},
		
			{loc:402653184, data: locIndex7500LOC402653184},
		
			{loc:671088640, data: locIndex7500LOC671088640},
		
			{loc:939524096, data: locIndex7500LOC939524096},
		
			{loc:1207959552, data: locIndex7500LOC1207959552},
		
			{loc:2147483648, data: locIndex7500LOC2147483648},
		
			{loc:2281701376, data: locIndex7500LOC2281701376},
		
			{loc:3355443200, data: locIndex7500LOC3355443200},
		
			{loc:3623878656, data: locIndex7500LOC3623878656},
		
			{loc:4294967296, data: locIndex7500LOC4294967296},
		
			{loc:6442450944, data: locIndex7500LOC6442450944},
		
			{loc:6576668672, data: locIndex7500LOC6576668672},
		
	}}
}
