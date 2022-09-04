package pkg

// Shared functions and other effects.

// Trim horizontal white space from start and end of utf-8 string.
func TrimWS(s string) string {
	var (
		i, j, n  uint32
		runes    []rune
		iCh, jCh chan uint32
		toCheck  = [4]rune{' ', '\n', '\t', '\r'}
	)

	// Inst. objects.
	runes = []rune(s)
	iCh, jCh = make(chan uint32), make(chan uint32)
	n = uint32(len(runes))

	// Check whitespace function.
	isWS := func(r rune) bool {
		for _, ws := range toCheck {
			if r == ws {
				return true
			}
		}
		return false
	}

	// Iterate from start and end. if whitespace character,
	// increment / decrement variable until false.
	fn := func(index uint32, outCh chan<- uint32, increment bool) {
		for isWS(runes[index]) {
			if increment {
				index++
			} else {
				index--
			}
		}
		outCh <- index		
	}
	i, j = 0, n-1
	go fn(i, iCh, true)
	go fn(j, jCh, false)

	return s[<-iCh : <-jCh+1]
}

// Transform const string into *string
func ConstToRef(s string) *string {
	var outs *string
	inRunes := []rune(s)
	*outs = string(inRunes)
	return outs
}
