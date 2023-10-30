package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_9500_1024.bin
	locIndex9500LOC1024 []byte
	
	//go:embed fs_data_9500_37564416.bin
	locIndex9500LOC37564416 []byte
	
	//go:embed fs_data_9500_134217728.bin
	locIndex9500LOC134217728 []byte
	
	//go:embed fs_data_9500_402653184.bin
	locIndex9500LOC402653184 []byte
	
	//go:embed fs_data_9500_671088640.bin
	locIndex9500LOC671088640 []byte
	
	//go:embed fs_data_9500_939524096.bin
	locIndex9500LOC939524096 []byte
	
	//go:embed fs_data_9500_1207959552.bin
	locIndex9500LOC1207959552 []byte
	
	//go:embed fs_data_9500_2147483648.bin
	locIndex9500LOC2147483648 []byte
	
	//go:embed fs_data_9500_3355443200.bin
	locIndex9500LOC3355443200 []byte
	
	//go:embed fs_data_9500_3623878656.bin
	locIndex9500LOC3623878656 []byte
	
	//go:embed fs_data_9500_4294967296.bin
	locIndex9500LOC4294967296 []byte
	
	//go:embed fs_data_9500_4429185024.bin
	locIndex9500LOC4429185024 []byte
	
	//go:embed fs_data_9500_6442450944.bin
	locIndex9500LOC6442450944 []byte
	
	//go:embed fs_data_9500_6576668672.bin
	locIndex9500LOC6576668672 []byte
	
	//go:embed fs_data_9500_8589934592.bin
	locIndex9500LOC8589934592 []byte
	
)

func init() {
 	allFileData[9500] = fileData{size:9961472000, data:[]dataAt{
		
			{loc:1024, data: locIndex9500LOC1024},
		
			{loc:37564416, data: locIndex9500LOC37564416},
		
			{loc:134217728, data: locIndex9500LOC134217728},
		
			{loc:402653184, data: locIndex9500LOC402653184},
		
			{loc:671088640, data: locIndex9500LOC671088640},
		
			{loc:939524096, data: locIndex9500LOC939524096},
		
			{loc:1207959552, data: locIndex9500LOC1207959552},
		
			{loc:2147483648, data: locIndex9500LOC2147483648},
		
			{loc:3355443200, data: locIndex9500LOC3355443200},
		
			{loc:3623878656, data: locIndex9500LOC3623878656},
		
			{loc:4294967296, data: locIndex9500LOC4294967296},
		
			{loc:4429185024, data: locIndex9500LOC4429185024},
		
			{loc:6442450944, data: locIndex9500LOC6442450944},
		
			{loc:6576668672, data: locIndex9500LOC6576668672},
		
			{loc:8589934592, data: locIndex9500LOC8589934592},
		
	}}
}
