package common

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ItemExists(arrayType []string, item string) bool {
	for _, typ := range arrayType {
		if typ == item {
			return true
		}
	}
	return false
}
