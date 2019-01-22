package stringutil

func Split(slice string, letters string) []string {
	rawSplit := byteutil.Split([]byte(slice), []byte(letters))
	result := []string{}
	for _, v := range rawSplit {
		result := append(result, string(v))
	}
	return result
}
