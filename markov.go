package main

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mb-14/gomarkov"
)

// Struct which impliments ServeHttp to associate
// data with each api endpoint.
// See http.Handler for more detail
type chainHandler struct {
	// Name/filepath of corpi
	corpi string
	chain *gomarkov.Chain
	// interface which tokenizes input file
	tokenizer   Tokenizer
	premutator  Premutator
	postmutator Postmutator
}

// Support for adding a new entry to the chain,
// including saving the file to disk for later
func (h *chainHandler) writeNewEntry(data []byte) error {
	hash := sha256.Sum256(data)
	log.Println("Writing new entry to:", string(hash[0:32]))
	// Add new data to chain
	h.chain.Add(h.tokenizer.tokenize(data))
	err := ioutil.WriteFile(filepath.Join(h.corpi, string(hash[0:32])), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Returns markov generated input
func (h *chainHandler) getRawInput() ([]string, error) {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, err := h.chain.Generate(tokens[(len(tokens) - 1):])
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, next)
	}
	return tokens, nil
}

// Postmutate data
func (h *chainHandler) getMutatedInput() ([]string, error) {
	data, err := h.getRawInput()
	return h.postmutator.Mutate(data), err
}

// Creates a new chain based on the data in fuzzingDir
// which is tokenized by a tokenzier
func createChain(corpus string, t Tokenizer, pre Premutator, post Postmutator) *chainHandler {
	chain := gomarkov.NewChain(1) // TODO figure out what this number is
	h := &chainHandler{corpus, chain, t, pre, post}
	h.loadDirectory(corpus)
	return h
}

// loadDirectory takes a directory and parses each file in it, applying
// the function tokenize to the filedata and adding the response to the
// chain instance
func (h *chainHandler) loadDirectory(corpus string) {
	// corpi does no authentication, should probably limit to current dir
	err := filepath.Walk(corpus, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // error accessing file, TODO test
		}

		// Ignore loading directories
		if info.IsDir() {
			return nil
		}

		log.Println("Parsing: ", path)
		data, err := h.readFile(path)

		// This is ineffecient, TODO make more efficent
		tokens := h.tokenizer.tokenize(data)
		h.chain.Add(tokens)
		return nil
	})
	if err != nil {
		log.Println("Error parsing file:", err) // err parsing file, print and ignore
	}
}

// Readfile wraps ioutil to mutate file
func (h *chainHandler) readFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	// Mutate data before it's tokenized
	return h.premutator.Mutate(data), err
}
