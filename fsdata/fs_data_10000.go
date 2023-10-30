package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_10000_1024.bin
	locIndex10000LOC1024 []byte
	
	//go:embed fs_data_10000_37564416.bin
	locIndex10000LOC37564416 []byte
	
	//go:embed fs_data_10000_134217728.bin
	locIndex10000LOC134217728 []byte
	
	//go:embed fs_data_10000_402653184.bin
	locIndex10000LOC402653184 []byte
	
	//go:embed fs_data_10000_671088640.bin
	locIndex10000LOC671088640 []byte
	
	//go:embed fs_data_10000_939524096.bin
	locIndex10000LOC939524096 []byte
	
	//go:embed fs_data_10000_1207959552.bin
	locIndex10000LOC1207959552 []byte
	
	//go:embed fs_data_10000_2147483648.bin
	locIndex10000LOC2147483648 []byte
	
	//go:embed fs_data_10000_3355443200.bin
	locIndex10000LOC3355443200 []byte
	
	//go:embed fs_data_10000_3623878656.bin
	locIndex10000LOC3623878656 []byte
	
	//go:embed fs_data_10000_4294967296.bin
	locIndex10000LOC4294967296 []byte
	
	//go:embed fs_data_10000_4429185024.bin
	locIndex10000LOC4429185024 []byte
	
	//go:embed fs_data_10000_6442450944.bin
	locIndex10000LOC6442450944 []byte
	
	//go:embed fs_data_10000_6576668672.bin
	locIndex10000LOC6576668672 []byte
	
	//go:embed fs_data_10000_8589934592.bin
	locIndex10000LOC8589934592 []byte
	
)

func init() {
 	allFileData[10000] = fileData{size:10485760000, data:[]dataAt{
		
			{loc:1024, data: locIndex10000LOC1024},
		
			{loc:37564416, data: locIndex10000LOC37564416},
		
			{loc:134217728, data: locIndex10000LOC134217728},
		
			{loc:402653184, data: locIndex10000LOC402653184},
		
			{loc:671088640, data: locIndex10000LOC671088640},
		
			{loc:939524096, data: locIndex10000LOC939524096},
		
			{loc:1207959552, data: locIndex10000LOC1207959552},
		
			{loc:2147483648, data: locIndex10000LOC2147483648},
		
			{loc:3355443200, data: locIndex10000LOC3355443200},
		
			{loc:3623878656, data: locIndex10000LOC3623878656},
		
			{loc:4294967296, data: locIndex10000LOC4294967296},
		
			{loc:4429185024, data: locIndex10000LOC4429185024},
		
			{loc:6442450944, data: locIndex10000LOC6442450944},
		
			{loc:6576668672, data: locIndex10000LOC6576668672},
		
			{loc:8589934592, data: locIndex10000LOC8589934592},
		
	}}
}
