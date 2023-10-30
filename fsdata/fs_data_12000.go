package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_12000_1024.bin
	locIndex12000LOC1024 []byte
	
	//go:embed fs_data_12000_37826560.bin
	locIndex12000LOC37826560 []byte
	
	//go:embed fs_data_12000_134217728.bin
	locIndex12000LOC134217728 []byte
	
	//go:embed fs_data_12000_402653184.bin
	locIndex12000LOC402653184 []byte
	
	//go:embed fs_data_12000_671088640.bin
	locIndex12000LOC671088640 []byte
	
	//go:embed fs_data_12000_939524096.bin
	locIndex12000LOC939524096 []byte
	
	//go:embed fs_data_12000_1207959552.bin
	locIndex12000LOC1207959552 []byte
	
	//go:embed fs_data_12000_2147483648.bin
	locIndex12000LOC2147483648 []byte
	
	//go:embed fs_data_12000_3355443200.bin
	locIndex12000LOC3355443200 []byte
	
	//go:embed fs_data_12000_3623878656.bin
	locIndex12000LOC3623878656 []byte
	
	//go:embed fs_data_12000_4294967296.bin
	locIndex12000LOC4294967296 []byte
	
	//go:embed fs_data_12000_4429185024.bin
	locIndex12000LOC4429185024 []byte
	
	//go:embed fs_data_12000_6442450944.bin
	locIndex12000LOC6442450944 []byte
	
	//go:embed fs_data_12000_6576668672.bin
	locIndex12000LOC6576668672 []byte
	
	//go:embed fs_data_12000_8589934592.bin
	locIndex12000LOC8589934592 []byte
	
	//go:embed fs_data_12000_10737418240.bin
	locIndex12000LOC10737418240 []byte
	
	//go:embed fs_data_12000_10871635968.bin
	locIndex12000LOC10871635968 []byte
	
)

func init() {
 	allFileData[12000] = fileData{size:12582912000, data:[]dataAt{
		
			{loc:1024, data: locIndex12000LOC1024},
		
			{loc:37826560, data: locIndex12000LOC37826560},
		
			{loc:134217728, data: locIndex12000LOC134217728},
		
			{loc:402653184, data: locIndex12000LOC402653184},
		
			{loc:671088640, data: locIndex12000LOC671088640},
		
			{loc:939524096, data: locIndex12000LOC939524096},
		
			{loc:1207959552, data: locIndex12000LOC1207959552},
		
			{loc:2147483648, data: locIndex12000LOC2147483648},
		
			{loc:3355443200, data: locIndex12000LOC3355443200},
		
			{loc:3623878656, data: locIndex12000LOC3623878656},
		
			{loc:4294967296, data: locIndex12000LOC4294967296},
		
			{loc:4429185024, data: locIndex12000LOC4429185024},
		
			{loc:6442450944, data: locIndex12000LOC6442450944},
		
			{loc:6576668672, data: locIndex12000LOC6576668672},
		
			{loc:8589934592, data: locIndex12000LOC8589934592},
		
			{loc:10737418240, data: locIndex12000LOC10737418240},
		
			{loc:10871635968, data: locIndex12000LOC10871635968},
		
	}}
}
