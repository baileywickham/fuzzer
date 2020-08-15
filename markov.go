package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/mb-14/gomarkov"
)

// Seed for fetching random elts of slice
// for markov generation

type chainHandler struct {
	// Name/filepath of corpi
	corpi string
	chain *gomarkov.Chain
	// Chain requires a seed to start generation
	seeds []string
}

// Returns a len 3 seed for the random generator
func (c *chainHandler) getSeed() []string {
	s := make([]string, 0)
	for i := 0; i < 3; i++ {
		s = append(s, c.seeds[rand.Intn(len(c.seeds))])
	}
	return s
}

// Creates a new chain based on the data in fuzzingDir
// which is tokenized by a tokenzier
func createChain(corpi string, t Tokenizer) *chainHandler {
	chain := gomarkov.NewChain(1) // TODO figure out what this number is

	seeds := loadDirectory(corpi, chain, t)

	return &chainHandler{corpi, chain, seeds}
}

// loadDirectory takes a directory and parses each file in it, applying
// the function tokenize to the filedata and adding the response to the
// chain instance
// it returns a slice of strings which is the seed for mk generation
func loadDirectory(fuzzingDir string, chain *gomarkov.Chain, t Tokenizer) []string {
	s := make([]string, 0)

	// fuzzingDir does no authentication, should probably limit to current dir
	err := filepath.Walk(fuzzingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // error accessing file, TODO test
		}

		// Ignore loading directories
		if info.IsDir() {
			return nil
		}

		log.Println("Parsing: ", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// This is ineffecient, TODO make more efficent
		tokens := t.tokenize(data)
		if len(tokens) < 4 {
			// If too few tokens, add all of them
			s = append(s, tokens...)
		} else {
			for i := 0; i < 2; i++ {
				// Add 2 tokens per file
				s = append(s, tokens[rand.Int()%len(tokens)])
			}
		}
		chain.Add(tokens)
		return nil
	})
	if err != nil {
		log.Println("Error parsing file:", err) // err parsing file, print and ignore
	}
	return s
}
