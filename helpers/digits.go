package helpers

func IsDigit(c string) bool {
	if len(c) != 1 {
		return false
	}

	return c >= "0" && c <= "9"
}
