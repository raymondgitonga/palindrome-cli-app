package internal

func IsPalindrome(str string) bool {
	startPointer := 0
	endPointer := len(str) - 1

	for startPointer < endPointer {
		if str[startPointer] == str[endPointer] {
			startPointer++
			endPointer--
		} else {
			return false
		}
	}
	return true
}
