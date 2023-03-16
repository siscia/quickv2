package quick

import (
	"testing"
	"unicode"
)

func TestASCIIString(t *testing.T) {
	f := func(s ASCIIString) bool {
		for _, c := range s {
			if c < 32 {
				return false
			}
			if c > 126 {
				return false
			}
		}
		return true
	}
	if err := Check(f, &Config{MaxCount: 10_000}); err != nil {
		t.Error(err)
	}
}

func TestUTF8String(t *testing.T) {
	f := func(s UTF8String) bool {
		for _, c := range s {
			if !unicode.IsGraphic(rune(c)) {
				return false
			}
		}
		return true
	}
	// this is way too slow
	if err := Check(f, &Config{MaxCount: 1}); err != nil {
		t.Error(err)
	}
}
