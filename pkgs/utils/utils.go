package utils

func charRange(min uint8, max uint8) []uint8 {
	if max < min {
		return []uint8{}
	}
	s := make([]uint8, max-min)
	for i := 0 ; min < max ; i++{
		s[i] = min
		min++
	}
	return s
}

func MakePermutations() []string {
	upperCaseCh := charRange('A','Z'+1)
	lowerCaseCh := charRange('a','z'+1)
	CharSet := make([]string,0, 26*26*2)
	for _, v := range upperCaseCh {
		for  _,v2 := range upperCaseCh {
			CharSet = append(CharSet,(string([]byte{byte(v), byte(v2)})))
		}
	}
	for _, v := range upperCaseCh {
		for  _,v2 := range lowerCaseCh {
			CharSet = append(CharSet,(string([]byte{byte(v), byte(v2)})))
		}
	}
	return CharSet
}
