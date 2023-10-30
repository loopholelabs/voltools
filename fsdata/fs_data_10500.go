package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_10500_1024.bin
	locIndex10500LOC1024 []byte
	
	//go:embed fs_data_10500_37564416.bin
	locIndex10500LOC37564416 []byte
	
	//go:embed fs_data_10500_134217728.bin
	locIndex10500LOC134217728 []byte
	
	//go:embed fs_data_10500_402653184.bin
	locIndex10500LOC402653184 []byte
	
	//go:embed fs_data_10500_671088640.bin
	locIndex10500LOC671088640 []byte
	
	//go:embed fs_data_10500_939524096.bin
	locIndex10500LOC939524096 []byte
	
	//go:embed fs_data_10500_1207959552.bin
	locIndex10500LOC1207959552 []byte
	
	//go:embed fs_data_10500_2147483648.bin
	locIndex10500LOC2147483648 []byte
	
	//go:embed fs_data_10500_3355443200.bin
	locIndex10500LOC3355443200 []byte
	
	//go:embed fs_data_10500_3623878656.bin
	locIndex10500LOC3623878656 []byte
	
	//go:embed fs_data_10500_4294967296.bin
	locIndex10500LOC4294967296 []byte
	
	//go:embed fs_data_10500_4429185024.bin
	locIndex10500LOC4429185024 []byte
	
	//go:embed fs_data_10500_6442450944.bin
	locIndex10500LOC6442450944 []byte
	
	//go:embed fs_data_10500_6576668672.bin
	locIndex10500LOC6576668672 []byte
	
	//go:embed fs_data_10500_8589934592.bin
	locIndex10500LOC8589934592 []byte
	
	//go:embed fs_data_10500_10737418240.bin
	locIndex10500LOC10737418240 []byte
	
	//go:embed fs_data_10500_10871635968.bin
	locIndex10500LOC10871635968 []byte
	
)

func init() {
 	allFileData[10500] = fileData{size:11010048000, data:[]dataAt{
		
			{loc:1024, data: locIndex10500LOC1024},
		
			{loc:37564416, data: locIndex10500LOC37564416},
		
			{loc:134217728, data: locIndex10500LOC134217728},
		
			{loc:402653184, data: locIndex10500LOC402653184},
		
			{loc:671088640, data: locIndex10500LOC671088640},
		
			{loc:939524096, data: locIndex10500LOC939524096},
		
			{loc:1207959552, data: locIndex10500LOC1207959552},
		
			{loc:2147483648, data: locIndex10500LOC2147483648},
		
			{loc:3355443200, data: locIndex10500LOC3355443200},
		
			{loc:3623878656, data: locIndex10500LOC3623878656},
		
			{loc:4294967296, data: locIndex10500LOC4294967296},
		
			{loc:4429185024, data: locIndex10500LOC4429185024},
		
			{loc:6442450944, data: locIndex10500LOC6442450944},
		
			{loc:6576668672, data: locIndex10500LOC6576668672},
		
			{loc:8589934592, data: locIndex10500LOC8589934592},
		
			{loc:10737418240, data: locIndex10500LOC10737418240},
		
			{loc:10871635968, data: locIndex10500LOC10871635968},
		
	}}
}
