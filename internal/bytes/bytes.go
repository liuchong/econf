package bytes

func IsLowercase(c uint8) bool {
	return c >= 'a' && c <= 'z'
}

func IsUppercase(c uint8) bool {
	return c >= 'A' && c <= 'Z'
}

func IsNumber(c uint8) bool {
	return c >= '0' && c <= '9'
}
