package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_8500_1024.bin
	locIndex8500LOC1024 []byte
	
	//go:embed fs_data_8500_37629952.bin
	locIndex8500LOC37629952 []byte
	
	//go:embed fs_data_8500_134217728.bin
	locIndex8500LOC134217728 []byte
	
	//go:embed fs_data_8500_402653184.bin
	locIndex8500LOC402653184 []byte
	
	//go:embed fs_data_8500_671088640.bin
	locIndex8500LOC671088640 []byte
	
	//go:embed fs_data_8500_939524096.bin
	locIndex8500LOC939524096 []byte
	
	//go:embed fs_data_8500_1207959552.bin
	locIndex8500LOC1207959552 []byte
	
	//go:embed fs_data_8500_2147483648.bin
	locIndex8500LOC2147483648 []byte
	
	//go:embed fs_data_8500_3355443200.bin
	locIndex8500LOC3355443200 []byte
	
	//go:embed fs_data_8500_3623878656.bin
	locIndex8500LOC3623878656 []byte
	
	//go:embed fs_data_8500_4294967296.bin
	locIndex8500LOC4294967296 []byte
	
	//go:embed fs_data_8500_4429185024.bin
	locIndex8500LOC4429185024 []byte
	
	//go:embed fs_data_8500_6442450944.bin
	locIndex8500LOC6442450944 []byte
	
	//go:embed fs_data_8500_6576668672.bin
	locIndex8500LOC6576668672 []byte
	
	//go:embed fs_data_8500_8589934592.bin
	locIndex8500LOC8589934592 []byte
	
)

func init() {
 	allFileData[8500] = fileData{size:8912896000, data:[]dataAt{
		
			{loc:1024, data: locIndex8500LOC1024},
		
			{loc:37629952, data: locIndex8500LOC37629952},
		
			{loc:134217728, data: locIndex8500LOC134217728},
		
			{loc:402653184, data: locIndex8500LOC402653184},
		
			{loc:671088640, data: locIndex8500LOC671088640},
		
			{loc:939524096, data: locIndex8500LOC939524096},
		
			{loc:1207959552, data: locIndex8500LOC1207959552},
		
			{loc:2147483648, data: locIndex8500LOC2147483648},
		
			{loc:3355443200, data: locIndex8500LOC3355443200},
		
			{loc:3623878656, data: locIndex8500LOC3623878656},
		
			{loc:4294967296, data: locIndex8500LOC4294967296},
		
			{loc:4429185024, data: locIndex8500LOC4429185024},
		
			{loc:6442450944, data: locIndex8500LOC6442450944},
		
			{loc:6576668672, data: locIndex8500LOC6576668672},
		
			{loc:8589934592, data: locIndex8500LOC8589934592},
		
	}}
}
