package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_11500_1025.bin
	locIndex11500LOC1025 []byte
	
	//go:embed fs_data_11500_37892096.bin
	locIndex11500LOC37892096 []byte
	
	//go:embed fs_data_11500_134217729.bin
	locIndex11500LOC134217729 []byte
	
	//go:embed fs_data_11500_402653185.bin
	locIndex11500LOC402653185 []byte
	
	//go:embed fs_data_11500_671088641.bin
	locIndex11500LOC671088641 []byte
	
	//go:embed fs_data_11500_939524097.bin
	locIndex11500LOC939524097 []byte
	
	//go:embed fs_data_11500_1207959553.bin
	locIndex11500LOC1207959553 []byte
	
	//go:embed fs_data_11500_2147483648.bin
	locIndex11500LOC2147483648 []byte
	
	//go:embed fs_data_11500_3355443201.bin
	locIndex11500LOC3355443201 []byte
	
	//go:embed fs_data_11500_3623878657.bin
	locIndex11500LOC3623878657 []byte
	
	//go:embed fs_data_11500_4294967296.bin
	locIndex11500LOC4294967296 []byte
	
	//go:embed fs_data_11500_4429185024.bin
	locIndex11500LOC4429185024 []byte
	
	//go:embed fs_data_11500_6442450944.bin
	locIndex11500LOC6442450944 []byte
	
	//go:embed fs_data_11500_6576668673.bin
	locIndex11500LOC6576668673 []byte
	
	//go:embed fs_data_11500_8589934592.bin
	locIndex11500LOC8589934592 []byte
	
	//go:embed fs_data_11500_10737418240.bin
	locIndex11500LOC10737418240 []byte
	
	//go:embed fs_data_11500_10871635969.bin
	locIndex11500LOC10871635969 []byte
	
)

func init() {
 	allFileData[11500] = fileData{size:12058624000, data:[]dataAt{
		
			{loc:1025, data: locIndex11500LOC1025},
		
			{loc:37892096, data: locIndex11500LOC37892096},
		
			{loc:134217729, data: locIndex11500LOC134217729},
		
			{loc:402653185, data: locIndex11500LOC402653185},
		
			{loc:671088641, data: locIndex11500LOC671088641},
		
			{loc:939524097, data: locIndex11500LOC939524097},
		
			{loc:1207959553, data: locIndex11500LOC1207959553},
		
			{loc:2147483648, data: locIndex11500LOC2147483648},
		
			{loc:3355443201, data: locIndex11500LOC3355443201},
		
			{loc:3623878657, data: locIndex11500LOC3623878657},
		
			{loc:4294967296, data: locIndex11500LOC4294967296},
		
			{loc:4429185024, data: locIndex11500LOC4429185024},
		
			{loc:6442450944, data: locIndex11500LOC6442450944},
		
			{loc:6576668673, data: locIndex11500LOC6576668673},
		
			{loc:8589934592, data: locIndex11500LOC8589934592},
		
			{loc:10737418240, data: locIndex11500LOC10737418240},
		
			{loc:10871635969, data: locIndex11500LOC10871635969},
		
	}}
}
