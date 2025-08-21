package util

func StringIfNotEmpty(newVal, oldVal string) string {
	if newVal != "" {
		return newVal
	}
	return oldVal
}

func JsonIfNotEmpty(newVal, oldVal []byte) []byte {
	if len(newVal) > 0 {
		return newVal
	}
	return oldVal
}
