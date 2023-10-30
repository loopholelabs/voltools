package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_13000_1024.bin
	locIndex13000LOC1024 []byte
	
	//go:embed fs_data_13000_37761024.bin
	locIndex13000LOC37761024 []byte
	
	//go:embed fs_data_13000_134217728.bin
	locIndex13000LOC134217728 []byte
	
	//go:embed fs_data_13000_402653184.bin
	locIndex13000LOC402653184 []byte
	
	//go:embed fs_data_13000_671088640.bin
	locIndex13000LOC671088640 []byte
	
	//go:embed fs_data_13000_939524096.bin
	locIndex13000LOC939524096 []byte
	
	//go:embed fs_data_13000_1207959552.bin
	locIndex13000LOC1207959552 []byte
	
	//go:embed fs_data_13000_2147483648.bin
	locIndex13000LOC2147483648 []byte
	
	//go:embed fs_data_13000_3355443200.bin
	locIndex13000LOC3355443200 []byte
	
	//go:embed fs_data_13000_3623878656.bin
	locIndex13000LOC3623878656 []byte
	
	//go:embed fs_data_13000_4294967296.bin
	locIndex13000LOC4294967296 []byte
	
	//go:embed fs_data_13000_6442450944.bin
	locIndex13000LOC6442450944 []byte
	
	//go:embed fs_data_13000_6576668672.bin
	locIndex13000LOC6576668672 []byte
	
	//go:embed fs_data_13000_6580875264.bin
	locIndex13000LOC6580875264 []byte
	
	//go:embed fs_data_13000_8589934592.bin
	locIndex13000LOC8589934592 []byte
	
	//go:embed fs_data_13000_10737418240.bin
	locIndex13000LOC10737418240 []byte
	
	//go:embed fs_data_13000_10871635968.bin
	locIndex13000LOC10871635968 []byte
	
	//go:embed fs_data_13000_12884901888.bin
	locIndex13000LOC12884901888 []byte
	
)

func init() {
 	allFileData[13000] = fileData{size:13631488000, data:[]dataAt{
		
			{loc:1024, data: locIndex13000LOC1024},
		
			{loc:37761024, data: locIndex13000LOC37761024},
		
			{loc:134217728, data: locIndex13000LOC134217728},
		
			{loc:402653184, data: locIndex13000LOC402653184},
		
			{loc:671088640, data: locIndex13000LOC671088640},
		
			{loc:939524096, data: locIndex13000LOC939524096},
		
			{loc:1207959552, data: locIndex13000LOC1207959552},
		
			{loc:2147483648, data: locIndex13000LOC2147483648},
		
			{loc:3355443200, data: locIndex13000LOC3355443200},
		
			{loc:3623878656, data: locIndex13000LOC3623878656},
		
			{loc:4294967296, data: locIndex13000LOC4294967296},
		
			{loc:6442450944, data: locIndex13000LOC6442450944},
		
			{loc:6576668672, data: locIndex13000LOC6576668672},
		
			{loc:6580875264, data: locIndex13000LOC6580875264},
		
			{loc:8589934592, data: locIndex13000LOC8589934592},
		
			{loc:10737418240, data: locIndex13000LOC10737418240},
		
			{loc:10871635968, data: locIndex13000LOC10871635968},
		
			{loc:12884901888, data: locIndex13000LOC12884901888},
		
	}}
}
