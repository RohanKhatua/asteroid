package utils

// isLetter checks if a character is a letter or an underscore.
func IsLetter(ch byte) bool {
	// check if the character is a letter or an underscore
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if a character is a digit.
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
