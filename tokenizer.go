package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-audio/wav"
)

type Tokenizer interface {
	tokenize(*os.File) []string
}

// There is probably a better way to do this, but each tokenizer
// should be a function which conforms to the Tokenizer interface,
// which seems to require an empty struct
type tokenizeBySpaces struct {
}

type tokenizeWav struct {
}

func (t tokenizeBySpaces) tokenize(file *os.File) []string {
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Println(err)
	}
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

func (t tokenizeWav) tokenize(file *os.File) []string {
	s := make([]string, 0)
	d := wav.NewDecoder(file)
	for c, err := d.NextChunk(); err != nil; c, err = d.NextChunk() {
		if err != nil {
			return s
		}
		buff := make([]byte, c.Size)
		c.R.Read(buff)
		s = append(s, string(buff))
	}
	return s
}
