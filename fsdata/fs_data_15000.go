package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_15000_1024.bin
	locIndex15000LOC1024 []byte
	
	//go:embed fs_data_15000_37695488.bin
	locIndex15000LOC37695488 []byte
	
	//go:embed fs_data_15000_134217728.bin
	locIndex15000LOC134217728 []byte
	
	//go:embed fs_data_15000_402653184.bin
	locIndex15000LOC402653184 []byte
	
	//go:embed fs_data_15000_671088640.bin
	locIndex15000LOC671088640 []byte
	
	//go:embed fs_data_15000_939524096.bin
	locIndex15000LOC939524096 []byte
	
	//go:embed fs_data_15000_1207959552.bin
	locIndex15000LOC1207959552 []byte
	
	//go:embed fs_data_15000_2147483648.bin
	locIndex15000LOC2147483648 []byte
	
	//go:embed fs_data_15000_3355443200.bin
	locIndex15000LOC3355443200 []byte
	
	//go:embed fs_data_15000_3623878656.bin
	locIndex15000LOC3623878656 []byte
	
	//go:embed fs_data_15000_4294967296.bin
	locIndex15000LOC4294967296 []byte
	
	//go:embed fs_data_15000_6442450944.bin
	locIndex15000LOC6442450944 []byte
	
	//go:embed fs_data_15000_6576668672.bin
	locIndex15000LOC6576668672 []byte
	
	//go:embed fs_data_15000_6580875264.bin
	locIndex15000LOC6580875264 []byte
	
	//go:embed fs_data_15000_8589934592.bin
	locIndex15000LOC8589934592 []byte
	
	//go:embed fs_data_15000_10737418240.bin
	locIndex15000LOC10737418240 []byte
	
	//go:embed fs_data_15000_10871635968.bin
	locIndex15000LOC10871635968 []byte
	
	//go:embed fs_data_15000_12884901888.bin
	locIndex15000LOC12884901888 []byte
	
	//go:embed fs_data_15000_15032385536.bin
	locIndex15000LOC15032385536 []byte
	
)

func init() {
 	allFileData[15000] = fileData{size:15728640000, data:[]dataAt{
		
			{loc:1024, data: locIndex15000LOC1024},
		
			{loc:37695488, data: locIndex15000LOC37695488},
		
			{loc:134217728, data: locIndex15000LOC134217728},
		
			{loc:402653184, data: locIndex15000LOC402653184},
		
			{loc:671088640, data: locIndex15000LOC671088640},
		
			{loc:939524096, data: locIndex15000LOC939524096},
		
			{loc:1207959552, data: locIndex15000LOC1207959552},
		
			{loc:2147483648, data: locIndex15000LOC2147483648},
		
			{loc:3355443200, data: locIndex15000LOC3355443200},
		
			{loc:3623878656, data: locIndex15000LOC3623878656},
		
			{loc:4294967296, data: locIndex15000LOC4294967296},
		
			{loc:6442450944, data: locIndex15000LOC6442450944},
		
			{loc:6576668672, data: locIndex15000LOC6576668672},
		
			{loc:6580875264, data: locIndex15000LOC6580875264},
		
			{loc:8589934592, data: locIndex15000LOC8589934592},
		
			{loc:10737418240, data: locIndex15000LOC10737418240},
		
			{loc:10871635968, data: locIndex15000LOC10871635968},
		
			{loc:12884901888, data: locIndex15000LOC12884901888},
		
			{loc:15032385536, data: locIndex15000LOC15032385536},
		
	}}
}
