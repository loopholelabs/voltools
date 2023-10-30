package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_3500_1025.bin
	locIndex3500LOC1025 []byte
	
	//go:embed fs_data_3500_34697216.bin
	locIndex3500LOC34697216 []byte
	
	//go:embed fs_data_3500_134217729.bin
	locIndex3500LOC134217729 []byte
	
	//go:embed fs_data_3500_402653185.bin
	locIndex3500LOC402653185 []byte
	
	//go:embed fs_data_3500_671088641.bin
	locIndex3500LOC671088641 []byte
	
	//go:embed fs_data_3500_939524097.bin
	locIndex3500LOC939524097 []byte
	
	//go:embed fs_data_3500_1207959553.bin
	locIndex3500LOC1207959553 []byte
	
	//go:embed fs_data_3500_1610612736.bin
	locIndex3500LOC1610612736 []byte
	
	//go:embed fs_data_3500_2147483648.bin
	locIndex3500LOC2147483648 []byte
	
	//go:embed fs_data_3500_3355443201.bin
	locIndex3500LOC3355443201 []byte
	
	//go:embed fs_data_3500_3623878657.bin
	locIndex3500LOC3623878657 []byte
	
)

func init() {
 	allFileData[3500] = fileData{size:3670016000, data:[]dataAt{
		
			{loc:1025, data: locIndex3500LOC1025},
		
			{loc:34697216, data: locIndex3500LOC34697216},
		
			{loc:134217729, data: locIndex3500LOC134217729},
		
			{loc:402653185, data: locIndex3500LOC402653185},
		
			{loc:671088641, data: locIndex3500LOC671088641},
		
			{loc:939524097, data: locIndex3500LOC939524097},
		
			{loc:1207959553, data: locIndex3500LOC1207959553},
		
			{loc:1610612736, data: locIndex3500LOC1610612736},
		
			{loc:2147483648, data: locIndex3500LOC2147483648},
		
			{loc:3355443201, data: locIndex3500LOC3355443201},
		
			{loc:3623878657, data: locIndex3500LOC3623878657},
		
	}}
}
