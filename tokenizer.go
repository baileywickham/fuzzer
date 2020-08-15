package main

import "strings"

type Tokenizer interface {
	tokenize([]byte) []string
}

// There is probably a better way to do this, but each tokenizer
// should be a function which conforms to the Tokenizer interface,
// which seems to require an empty struct
type tokenizeBySpaces struct {
}

func (t tokenizeBySpaces) tokenize(data []byte) []string {
	return strings.Split(string(data), " ")
}
