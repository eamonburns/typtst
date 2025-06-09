package words

import (
	_ "embed"
	"strings"
)

//go:embed 1000en
var words1000en string

var Words = strings.Split(words1000en, "\n")
