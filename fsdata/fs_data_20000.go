package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_20000_1024.bin
	locIndex20000LOC1024 []byte
	
	//go:embed fs_data_20000_37765120.bin
	locIndex20000LOC37765120 []byte
	
	//go:embed fs_data_20000_134217728.bin
	locIndex20000LOC134217728 []byte
	
	//go:embed fs_data_20000_402653184.bin
	locIndex20000LOC402653184 []byte
	
	//go:embed fs_data_20000_671088640.bin
	locIndex20000LOC671088640 []byte
	
	//go:embed fs_data_20000_939524096.bin
	locIndex20000LOC939524096 []byte
	
	//go:embed fs_data_20000_1207959552.bin
	locIndex20000LOC1207959552 []byte
	
	//go:embed fs_data_20000_2147483648.bin
	locIndex20000LOC2147483648 []byte
	
	//go:embed fs_data_20000_3355443200.bin
	locIndex20000LOC3355443200 []byte
	
	//go:embed fs_data_20000_3623878656.bin
	locIndex20000LOC3623878656 []byte
	
	//go:embed fs_data_20000_4294967296.bin
	locIndex20000LOC4294967296 []byte
	
	//go:embed fs_data_20000_6442450944.bin
	locIndex20000LOC6442450944 []byte
	
	//go:embed fs_data_20000_6576668672.bin
	locIndex20000LOC6576668672 []byte
	
	//go:embed fs_data_20000_8589934592.bin
	locIndex20000LOC8589934592 []byte
	
	//go:embed fs_data_20000_8724152320.bin
	locIndex20000LOC8724152320 []byte
	
	//go:embed fs_data_20000_10737418240.bin
	locIndex20000LOC10737418240 []byte
	
	//go:embed fs_data_20000_10871635968.bin
	locIndex20000LOC10871635968 []byte
	
	//go:embed fs_data_20000_12884901888.bin
	locIndex20000LOC12884901888 []byte
	
	//go:embed fs_data_20000_15032385536.bin
	locIndex20000LOC15032385536 []byte
	
	//go:embed fs_data_20000_16777216000.bin
	locIndex20000LOC16777216000 []byte
	
	//go:embed fs_data_20000_17179869184.bin
	locIndex20000LOC17179869184 []byte
	
	//go:embed fs_data_20000_19327352832.bin
	locIndex20000LOC19327352832 []byte
	
)

func init() {
 	allFileData[20000] = fileData{size:20971520000, data:[]dataAt{
		
			{loc:1024, data: locIndex20000LOC1024},
		
			{loc:37765120, data: locIndex20000LOC37765120},
		
			{loc:134217728, data: locIndex20000LOC134217728},
		
			{loc:402653184, data: locIndex20000LOC402653184},
		
			{loc:671088640, data: locIndex20000LOC671088640},
		
			{loc:939524096, data: locIndex20000LOC939524096},
		
			{loc:1207959552, data: locIndex20000LOC1207959552},
		
			{loc:2147483648, data: locIndex20000LOC2147483648},
		
			{loc:3355443200, data: locIndex20000LOC3355443200},
		
			{loc:3623878656, data: locIndex20000LOC3623878656},
		
			{loc:4294967296, data: locIndex20000LOC4294967296},
		
			{loc:6442450944, data: locIndex20000LOC6442450944},
		
			{loc:6576668672, data: locIndex20000LOC6576668672},
		
			{loc:8589934592, data: locIndex20000LOC8589934592},
		
			{loc:8724152320, data: locIndex20000LOC8724152320},
		
			{loc:10737418240, data: locIndex20000LOC10737418240},
		
			{loc:10871635968, data: locIndex20000LOC10871635968},
		
			{loc:12884901888, data: locIndex20000LOC12884901888},
		
			{loc:15032385536, data: locIndex20000LOC15032385536},
		
			{loc:16777216000, data: locIndex20000LOC16777216000},
		
			{loc:17179869184, data: locIndex20000LOC17179869184},
		
			{loc:19327352832, data: locIndex20000LOC19327352832},
		
	}}
}
