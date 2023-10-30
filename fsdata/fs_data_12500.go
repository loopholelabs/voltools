package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_12500_1024.bin
	locIndex12500LOC1024 []byte
	
	//go:embed fs_data_12500_37826560.bin
	locIndex12500LOC37826560 []byte
	
	//go:embed fs_data_12500_134217728.bin
	locIndex12500LOC134217728 []byte
	
	//go:embed fs_data_12500_402653184.bin
	locIndex12500LOC402653184 []byte
	
	//go:embed fs_data_12500_671088640.bin
	locIndex12500LOC671088640 []byte
	
	//go:embed fs_data_12500_939524096.bin
	locIndex12500LOC939524096 []byte
	
	//go:embed fs_data_12500_1207959552.bin
	locIndex12500LOC1207959552 []byte
	
	//go:embed fs_data_12500_2147483648.bin
	locIndex12500LOC2147483648 []byte
	
	//go:embed fs_data_12500_3355443200.bin
	locIndex12500LOC3355443200 []byte
	
	//go:embed fs_data_12500_3623878656.bin
	locIndex12500LOC3623878656 []byte
	
	//go:embed fs_data_12500_4294967296.bin
	locIndex12500LOC4294967296 []byte
	
	//go:embed fs_data_12500_6442450944.bin
	locIndex12500LOC6442450944 []byte
	
	//go:embed fs_data_12500_6576668672.bin
	locIndex12500LOC6576668672 []byte
	
	//go:embed fs_data_12500_6580875264.bin
	locIndex12500LOC6580875264 []byte
	
	//go:embed fs_data_12500_8589934592.bin
	locIndex12500LOC8589934592 []byte
	
	//go:embed fs_data_12500_10737418240.bin
	locIndex12500LOC10737418240 []byte
	
	//go:embed fs_data_12500_10871635968.bin
	locIndex12500LOC10871635968 []byte
	
	//go:embed fs_data_12500_12884901888.bin
	locIndex12500LOC12884901888 []byte
	
)

func init() {
 	allFileData[12500] = fileData{size:13107200000, data:[]dataAt{
		
			{loc:1024, data: locIndex12500LOC1024},
		
			{loc:37826560, data: locIndex12500LOC37826560},
		
			{loc:134217728, data: locIndex12500LOC134217728},
		
			{loc:402653184, data: locIndex12500LOC402653184},
		
			{loc:671088640, data: locIndex12500LOC671088640},
		
			{loc:939524096, data: locIndex12500LOC939524096},
		
			{loc:1207959552, data: locIndex12500LOC1207959552},
		
			{loc:2147483648, data: locIndex12500LOC2147483648},
		
			{loc:3355443200, data: locIndex12500LOC3355443200},
		
			{loc:3623878656, data: locIndex12500LOC3623878656},
		
			{loc:4294967296, data: locIndex12500LOC4294967296},
		
			{loc:6442450944, data: locIndex12500LOC6442450944},
		
			{loc:6576668672, data: locIndex12500LOC6576668672},
		
			{loc:6580875264, data: locIndex12500LOC6580875264},
		
			{loc:8589934592, data: locIndex12500LOC8589934592},
		
			{loc:10737418240, data: locIndex12500LOC10737418240},
		
			{loc:10871635968, data: locIndex12500LOC10871635968},
		
			{loc:12884901888, data: locIndex12500LOC12884901888},
		
	}}
}
