package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_6500_1024.bin
	locIndex6500LOC1024 []byte
	
	//go:embed fs_data_6500_36888576.bin
	locIndex6500LOC36888576 []byte
	
	//go:embed fs_data_6500_134217728.bin
	locIndex6500LOC134217728 []byte
	
	//go:embed fs_data_6500_402653184.bin
	locIndex6500LOC402653184 []byte
	
	//go:embed fs_data_6500_671088640.bin
	locIndex6500LOC671088640 []byte
	
	//go:embed fs_data_6500_939524096.bin
	locIndex6500LOC939524096 []byte
	
	//go:embed fs_data_6500_1207959552.bin
	locIndex6500LOC1207959552 []byte
	
	//go:embed fs_data_6500_2147483648.bin
	locIndex6500LOC2147483648 []byte
	
	//go:embed fs_data_6500_2281701376.bin
	locIndex6500LOC2281701376 []byte
	
	//go:embed fs_data_6500_3355443200.bin
	locIndex6500LOC3355443200 []byte
	
	//go:embed fs_data_6500_3623878656.bin
	locIndex6500LOC3623878656 []byte
	
	//go:embed fs_data_6500_4294967296.bin
	locIndex6500LOC4294967296 []byte
	
	//go:embed fs_data_6500_6442450944.bin
	locIndex6500LOC6442450944 []byte
	
	//go:embed fs_data_6500_6576668672.bin
	locIndex6500LOC6576668672 []byte
	
)

func init() {
 	allFileData[6500] = fileData{size:6815744000, data:[]dataAt{
		
			{loc:1024, data: locIndex6500LOC1024},
		
			{loc:36888576, data: locIndex6500LOC36888576},
		
			{loc:134217728, data: locIndex6500LOC134217728},
		
			{loc:402653184, data: locIndex6500LOC402653184},
		
			{loc:671088640, data: locIndex6500LOC671088640},
		
			{loc:939524096, data: locIndex6500LOC939524096},
		
			{loc:1207959552, data: locIndex6500LOC1207959552},
		
			{loc:2147483648, data: locIndex6500LOC2147483648},
		
			{loc:2281701376, data: locIndex6500LOC2281701376},
		
			{loc:3355443200, data: locIndex6500LOC3355443200},
		
			{loc:3623878656, data: locIndex6500LOC3623878656},
		
			{loc:4294967296, data: locIndex6500LOC4294967296},
		
			{loc:6442450944, data: locIndex6500LOC6442450944},
		
			{loc:6576668672, data: locIndex6500LOC6576668672},
		
	}}
}
