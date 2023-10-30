package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_1500_1025.bin
	locIndex1500LOC1025 []byte
	
	//go:embed fs_data_1500_25448448.bin
	locIndex1500LOC25448448 []byte
	
	//go:embed fs_data_1500_134217729.bin
	locIndex1500LOC134217729 []byte
	
	//go:embed fs_data_1500_402653185.bin
	locIndex1500LOC402653185 []byte
	
	//go:embed fs_data_1500_536870912.bin
	locIndex1500LOC536870912 []byte
	
	//go:embed fs_data_1500_671088641.bin
	locIndex1500LOC671088641 []byte
	
	//go:embed fs_data_1500_939524097.bin
	locIndex1500LOC939524097 []byte
	
	//go:embed fs_data_1500_1207959553.bin
	locIndex1500LOC1207959553 []byte
	
)

func init() {
 	allFileData[1500] = fileData{size:1572864000, data:[]dataAt{
		
			{loc:1025, data: locIndex1500LOC1025},
		
			{loc:25448448, data: locIndex1500LOC25448448},
		
			{loc:134217729, data: locIndex1500LOC134217729},
		
			{loc:402653185, data: locIndex1500LOC402653185},
		
			{loc:536870912, data: locIndex1500LOC536870912},
		
			{loc:671088641, data: locIndex1500LOC671088641},
		
			{loc:939524097, data: locIndex1500LOC939524097},
		
			{loc:1207959553, data: locIndex1500LOC1207959553},
		
	}}
}
