package main

import (
	"strings"
)

type Tokenizer interface {
	tokenize([]byte) []string
}

// There is probably a better way to do this, but each tokenizer
// should be a function which conforms to the Tokenizer interface,
// which seems to require an empty struct
type tokenizeBySpaces struct {
}

func (t tokenizeBySpaces) tokenize(data []byte) []string {
	strs := strings.Split(string(data), " ")
	s := make([]string, 0)
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}
		s = append(s, str)
	}
	return s
}
