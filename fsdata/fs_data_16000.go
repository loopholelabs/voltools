package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_16000_1025.bin
	locIndex16000LOC1025 []byte
	
	//go:embed fs_data_16000_37892096.bin
	locIndex16000LOC37892096 []byte
	
	//go:embed fs_data_16000_134217729.bin
	locIndex16000LOC134217729 []byte
	
	//go:embed fs_data_16000_402653185.bin
	locIndex16000LOC402653185 []byte
	
	//go:embed fs_data_16000_671088641.bin
	locIndex16000LOC671088641 []byte
	
	//go:embed fs_data_16000_939524097.bin
	locIndex16000LOC939524097 []byte
	
	//go:embed fs_data_16000_1207959553.bin
	locIndex16000LOC1207959553 []byte
	
	//go:embed fs_data_16000_2147483648.bin
	locIndex16000LOC2147483648 []byte
	
	//go:embed fs_data_16000_3355443201.bin
	locIndex16000LOC3355443201 []byte
	
	//go:embed fs_data_16000_3623878657.bin
	locIndex16000LOC3623878657 []byte
	
	//go:embed fs_data_16000_4294967296.bin
	locIndex16000LOC4294967296 []byte
	
	//go:embed fs_data_16000_6442450944.bin
	locIndex16000LOC6442450944 []byte
	
	//go:embed fs_data_16000_6576668673.bin
	locIndex16000LOC6576668673 []byte
	
	//go:embed fs_data_16000_6580875264.bin
	locIndex16000LOC6580875264 []byte
	
	//go:embed fs_data_16000_8589934592.bin
	locIndex16000LOC8589934592 []byte
	
	//go:embed fs_data_16000_10737418240.bin
	locIndex16000LOC10737418240 []byte
	
	//go:embed fs_data_16000_10871635969.bin
	locIndex16000LOC10871635969 []byte
	
	//go:embed fs_data_16000_12884901888.bin
	locIndex16000LOC12884901888 []byte
	
	//go:embed fs_data_16000_15032385536.bin
	locIndex16000LOC15032385536 []byte
	
)

func init() {
 	allFileData[16000] = fileData{size:16777216000, data:[]dataAt{
		
			{loc:1025, data: locIndex16000LOC1025},
		
			{loc:37892096, data: locIndex16000LOC37892096},
		
			{loc:134217729, data: locIndex16000LOC134217729},
		
			{loc:402653185, data: locIndex16000LOC402653185},
		
			{loc:671088641, data: locIndex16000LOC671088641},
		
			{loc:939524097, data: locIndex16000LOC939524097},
		
			{loc:1207959553, data: locIndex16000LOC1207959553},
		
			{loc:2147483648, data: locIndex16000LOC2147483648},
		
			{loc:3355443201, data: locIndex16000LOC3355443201},
		
			{loc:3623878657, data: locIndex16000LOC3623878657},
		
			{loc:4294967296, data: locIndex16000LOC4294967296},
		
			{loc:6442450944, data: locIndex16000LOC6442450944},
		
			{loc:6576668673, data: locIndex16000LOC6576668673},
		
			{loc:6580875264, data: locIndex16000LOC6580875264},
		
			{loc:8589934592, data: locIndex16000LOC8589934592},
		
			{loc:10737418240, data: locIndex16000LOC10737418240},
		
			{loc:10871635969, data: locIndex16000LOC10871635969},
		
			{loc:12884901888, data: locIndex16000LOC12884901888},
		
			{loc:15032385536, data: locIndex16000LOC15032385536},
		
	}}
}
