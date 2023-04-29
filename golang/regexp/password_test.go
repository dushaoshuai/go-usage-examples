package regexp_test

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/dushaoshuai/goloop"
)

func TestPassword(t *testing.T) {
	validPasswd := regexp.MustCompile(`[[:graph:]]{6,16}`)

	// Generate a table contains chars from '!' to '~'.
	// See https://en.wikipedia.org/wiki/ASCII#Printable_characters
	// and https://pkg.go.dev/regexp/syntax@go1.20.3#:~:text=digits%20(%3D%3D%20%5B0%2D9%5D)%0A%5B%5B%3A-,graph,-%3A%5D%5D%20%20%20%20graphical%20(%3D%3D%20%5B!%2D~%5D%20%3D%3D%20%5BA%2DZa
	// for details.
	var graphs []byte
	for i := range goloop.Range('!', '~') {
		graphs = append(graphs, byte(i.I))
	}

	for range goloop.Repeat(10000) {
		t.Run("", func(t *testing.T) {
			// Generate a random password.
			rand.Shuffle(len(graphs), func(i, j int) {
				graphs[i], graphs[j] = graphs[j], graphs[i]
			})
			passwd := graphs[:rand.Intn(11)+6]

			if !validPasswd.Match(passwd) {
				t.Errorf("password %q doesn't match", passwd)
			}
		})
	}
}
