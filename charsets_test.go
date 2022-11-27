package charsets

import (
	"math/rand"
	"testing"
)

func TestStringWithCharset(t *testing.T) {
	for i := 1; i < 10; i++ {
		randomNumberOfRunes := rand.Intn(50)
		randomRunesASCII := RandomRunes(randomNumberOfRunes, []rune(ASCII))
		randomRunesUnicode := RandomRunes(randomNumberOfRunes, []rune(Unicode))

		if len(randomRunesASCII) != randomNumberOfRunes {
			t.Errorf("expected slice of ascii runes of lenght %d, got %d", randomNumberOfRunes, len(randomRunesASCII))
		}

		if len(randomRunesUnicode) != randomNumberOfRunes {
			t.Errorf("expected slice of unicode runes of lenght %d, got %d", randomNumberOfRunes, len(randomRunesUnicode))
		}
	}
}
