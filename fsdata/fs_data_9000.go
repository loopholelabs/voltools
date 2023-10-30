package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_9000_1024.bin
	locIndex9000LOC1024 []byte
	
	//go:embed fs_data_9000_37629952.bin
	locIndex9000LOC37629952 []byte
	
	//go:embed fs_data_9000_134217728.bin
	locIndex9000LOC134217728 []byte
	
	//go:embed fs_data_9000_402653184.bin
	locIndex9000LOC402653184 []byte
	
	//go:embed fs_data_9000_671088640.bin
	locIndex9000LOC671088640 []byte
	
	//go:embed fs_data_9000_939524096.bin
	locIndex9000LOC939524096 []byte
	
	//go:embed fs_data_9000_1207959552.bin
	locIndex9000LOC1207959552 []byte
	
	//go:embed fs_data_9000_2147483648.bin
	locIndex9000LOC2147483648 []byte
	
	//go:embed fs_data_9000_3355443200.bin
	locIndex9000LOC3355443200 []byte
	
	//go:embed fs_data_9000_3623878656.bin
	locIndex9000LOC3623878656 []byte
	
	//go:embed fs_data_9000_4294967296.bin
	locIndex9000LOC4294967296 []byte
	
	//go:embed fs_data_9000_4429185024.bin
	locIndex9000LOC4429185024 []byte
	
	//go:embed fs_data_9000_6442450944.bin
	locIndex9000LOC6442450944 []byte
	
	//go:embed fs_data_9000_6576668672.bin
	locIndex9000LOC6576668672 []byte
	
	//go:embed fs_data_9000_8589934592.bin
	locIndex9000LOC8589934592 []byte
	
)

func init() {
 	allFileData[9000] = fileData{size:9437184000, data:[]dataAt{
		
			{loc:1024, data: locIndex9000LOC1024},
		
			{loc:37629952, data: locIndex9000LOC37629952},
		
			{loc:134217728, data: locIndex9000LOC134217728},
		
			{loc:402653184, data: locIndex9000LOC402653184},
		
			{loc:671088640, data: locIndex9000LOC671088640},
		
			{loc:939524096, data: locIndex9000LOC939524096},
		
			{loc:1207959552, data: locIndex9000LOC1207959552},
		
			{loc:2147483648, data: locIndex9000LOC2147483648},
		
			{loc:3355443200, data: locIndex9000LOC3355443200},
		
			{loc:3623878656, data: locIndex9000LOC3623878656},
		
			{loc:4294967296, data: locIndex9000LOC4294967296},
		
			{loc:4429185024, data: locIndex9000LOC4429185024},
		
			{loc:6442450944, data: locIndex9000LOC6442450944},
		
			{loc:6576668672, data: locIndex9000LOC6576668672},
		
			{loc:8589934592, data: locIndex9000LOC8589934592},
		
	}}
}
