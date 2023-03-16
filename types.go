package quick

import (
	"math/rand"
	"reflect"
	"strings"
	"unicode"
)

type ASCIIString string

func (a ASCIIString) Generate(rand *rand.Rand, size int) reflect.Value {
	s := strings.Builder{}
	for i := 0; i < size; i++ {
		// ASCII code between 32 and 127 are printable ASCII
		// However value 127 represent the DELETE code, that we are not interested in.
		// But the interval of Intn is half open, and we must +1 to the last number.
		// In total we delete the code 127, but we add 1 to the last number, keeping
		// us with 127.
		n := rand.Intn(127-32) + 32
		s.WriteByte(byte(n))
	}
	ascii := ASCIIString(s.String())
	return reflect.ValueOf(ascii)
}

type UTF8String string

// this is way to slow - needs to figure out a better way
func (a UTF8String) Generate(rand *rand.Rand, size int) reflect.Value {
	s := strings.Builder{}
	i := size
	for i > 0 {
		c := rand.Int31()
		if unicode.IsGraphic(rune(c)) {
			s.WriteRune(rune(c))
			i--
		}
	}
	ascii := UTF8String(s.String())
	return reflect.ValueOf(ascii)
}
