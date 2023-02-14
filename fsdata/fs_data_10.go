package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_10_1026.bin
	locIndex10LOC1026 []byte
	
	//go:embed fs_data_10_37892096.bin
	locIndex10LOC37892096 []byte
	
	//go:embed fs_data_10_134217730.bin
	locIndex10LOC134217730 []byte
	
	//go:embed fs_data_10_402653186.bin
	locIndex10LOC402653186 []byte
	
	//go:embed fs_data_10_671088642.bin
	locIndex10LOC671088642 []byte
	
	//go:embed fs_data_10_939524098.bin
	locIndex10LOC939524098 []byte
	
	//go:embed fs_data_10_1207959554.bin
	locIndex10LOC1207959554 []byte
	
	//go:embed fs_data_10_2147483648.bin
	locIndex10LOC2147483648 []byte
	
	//go:embed fs_data_10_3355443202.bin
	locIndex10LOC3355443202 []byte
	
	//go:embed fs_data_10_3623878658.bin
	locIndex10LOC3623878658 []byte
	
	//go:embed fs_data_10_4294967296.bin
	locIndex10LOC4294967296 []byte
	
	//go:embed fs_data_10_4429185024.bin
	locIndex10LOC4429185024 []byte
	
	//go:embed fs_data_10_6442450944.bin
	locIndex10LOC6442450944 []byte
	
	//go:embed fs_data_10_6576668674.bin
	locIndex10LOC6576668674 []byte
	
	//go:embed fs_data_10_8589934592.bin
	locIndex10LOC8589934592 []byte
	
)

func init() {
 	allFileData[10] = fileData{size:10737418240, data:[]dataAt{
		
			{loc:1026, data: locIndex10LOC1026},
		
			{loc:37892096, data: locIndex10LOC37892096},
		
			{loc:134217730, data: locIndex10LOC134217730},
		
			{loc:402653186, data: locIndex10LOC402653186},
		
			{loc:671088642, data: locIndex10LOC671088642},
		
			{loc:939524098, data: locIndex10LOC939524098},
		
			{loc:1207959554, data: locIndex10LOC1207959554},
		
			{loc:2147483648, data: locIndex10LOC2147483648},
		
			{loc:3355443202, data: locIndex10LOC3355443202},
		
			{loc:3623878658, data: locIndex10LOC3623878658},
		
			{loc:4294967296, data: locIndex10LOC4294967296},
		
			{loc:4429185024, data: locIndex10LOC4429185024},
		
			{loc:6442450944, data: locIndex10LOC6442450944},
		
			{loc:6576668674, data: locIndex10LOC6576668674},
		
			{loc:8589934592, data: locIndex10LOC8589934592},
		
	}}
}
