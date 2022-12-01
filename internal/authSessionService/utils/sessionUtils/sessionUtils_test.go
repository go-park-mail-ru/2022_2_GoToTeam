package sessionUtils

import "testing"

func TestSessionRepository(t *testing.T) {
	l1 := GenerateRandomRunesString(111)
	l2 := GenerateRandomRunesString(111)

	if l1 == l2 {
		l1 = GenerateRandomRunesString(111)
		l2 = GenerateRandomRunesString(111)
		if l1 == l2 {
			t.Error("equals two times (")
		}
	}

}
