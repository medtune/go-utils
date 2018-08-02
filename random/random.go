package random

import (
	"math/rand"
	"time"

	"github.com/medtune/go-utils/errors"
)

var (
	alphaUp  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaLow = []rune("abcdefghijklmnopqrstuvwxyz")
	alpha    = []rune(append(alphaLow, alphaUp...)) // alphaLow + alphaUp
	numbers  = []rune("0123456789")
	special  = []rune("&()[]§!-_°*%^`$€=:;,?./+<>")
	alphaNum = []rune(append(alpha, numbers...)) // alpha + numbers
	mixin    = []rune(append(alphaNum, special...))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randFromRunes base string random generator from
// a runes array
func randFromRunes(runes []rune, lenght int) string {
	b := make([]rune, lenght)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// FromRunes return a random string from given runes array
func FromRunes(runes []rune, lenght int) (string, error) {
	if len(runes) == 0 {
		return "", errors.New("empty runes list")
	}
	return randFromRunes(runes, lenght), nil
}

// String return a totally random string
func String(lenght int) string {
	return randFromRunes(mixin, lenght)
}

// AlphaNumeric return a random string
// containing only alphabet chars and
// numericals
func AlphaNumeric(lenght int) string {
	return randFromRunes(alphaNum, lenght)
}

// Number return a random int between min
// and max
func Number(min, max int) int {
	return min + rand.Intn(max-min)
}

// Alphabetic return a random string containing
// only alphabetic characters (Uper and lower case)
func Alphabetic(lenght int) string {
	return randFromRunes(alpha, lenght)
}

// AlphaLower return a random string containing
// only alphabetic lowercase characters
func AlphaLower(lenght int) string {
	return randFromRunes(alphaLow, lenght)
}

// AlphaUpper return a random string containing
// only alphabetic uppercase characters
func AlphaUpper(lenght int) string {
	return randFromRunes(alphaUp, lenght)
}
