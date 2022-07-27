package validator

func StringsAreEqual(left, right string) bool {
	return left == right
}

func StringsAreDifferent(left, right string) bool {
	return left != right
}

func StringContainsDigit(inputText string) bool {
	for _, character := range inputText {
		if character >= '0' && character <= '9' {
			return true
		}
	}
	return false
}