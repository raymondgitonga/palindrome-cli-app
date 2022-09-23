package internal

type Palindrome struct {
	word string
}

func NewPalindrome(word string) *Palindrome {
	return &Palindrome{word: word}
}

func (p Palindrome) IsPalindrome() bool {
	startPointer := 0
	endPointer := len(p.word) - 1

	for startPointer < endPointer {
		if p.word[startPointer] == p.word[endPointer] {
			startPointer++
			endPointer--
		} else {
			return false
		}
	}
	return true
}
