package fsdata

import _ "embed"



var (
	
	//go:embed fs_data_500_1024.bin
	locIndex500LOC1024 []byte
	
	//go:embed fs_data_500_8103936.bin
	locIndex500LOC8103936 []byte
	
	//go:embed fs_data_500_25166848.bin
	locIndex500LOC25166848 []byte
	
	//go:embed fs_data_500_41944064.bin
	locIndex500LOC41944064 []byte
	
	//go:embed fs_data_500_58721280.bin
	locIndex500LOC58721280 []byte
	
	//go:embed fs_data_500_75498496.bin
	locIndex500LOC75498496 []byte
	
	//go:embed fs_data_500_134218752.bin
	locIndex500LOC134218752 []byte
	
	//go:embed fs_data_500_142607360.bin
	locIndex500LOC142607360 []byte
	
	//go:embed fs_data_500_209716224.bin
	locIndex500LOC209716224 []byte
	
	//go:embed fs_data_500_226493440.bin
	locIndex500LOC226493440 []byte
	
	//go:embed fs_data_500_268436480.bin
	locIndex500LOC268436480 []byte
	
	//go:embed fs_data_500_402654208.bin
	locIndex500LOC402654208 []byte
	
	//go:embed fs_data_500_411042816.bin
	locIndex500LOC411042816 []byte
	
)

func init() {
 	allFileData[500] = fileData{size:524288000, data:[]dataAt{
		
			{loc:1024, data: locIndex500LOC1024},
		
			{loc:8103936, data: locIndex500LOC8103936},
		
			{loc:25166848, data: locIndex500LOC25166848},
		
			{loc:41944064, data: locIndex500LOC41944064},
		
			{loc:58721280, data: locIndex500LOC58721280},
		
			{loc:75498496, data: locIndex500LOC75498496},
		
			{loc:134218752, data: locIndex500LOC134218752},
		
			{loc:142607360, data: locIndex500LOC142607360},
		
			{loc:209716224, data: locIndex500LOC209716224},
		
			{loc:226493440, data: locIndex500LOC226493440},
		
			{loc:268436480, data: locIndex500LOC268436480},
		
			{loc:402654208, data: locIndex500LOC402654208},
		
			{loc:411042816, data: locIndex500LOC411042816},
		
	}}
}
