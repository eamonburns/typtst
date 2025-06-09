package sentence

import (
	"math/rand"
	"strings"

	"github.com/agent-e11/typtst/internal/words"
)

func GenerateRandom(count int) string {
	s := strings.Builder{}

	for i := range count {
		if i != 0 {
			// Add spaces between words after the first iteration
			s.WriteString(" ")
		}
		s.WriteString(words.Words[rand.Intn(len(words.Words))])
	}

	return s.String()
}
