package controlflow

func OneComplement(bits string) string {
	comlement := ""
	for _, bit := range bits {
		if bit == '0' {
			comlement += "1"
		} else {
			comlement += "0"
		}
	}
	return comlement

}
