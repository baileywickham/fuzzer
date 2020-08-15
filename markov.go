package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mb-14/gomarkov"
)

func createChain(fuzzingDir string) {
	chain := gomarkov.NewChain(1) // TODO figure out what this number is
	loadDirectory(fuzzingDir, chain)
}

func getNextMarkovPred() (string, error) {
	return "hello", nil
}

func loadDirectory(fuzzingDir string, chain *gomarkov.Chain) {
	// fuzzingDir does no authentication, should probably limit to current dir
	err := filepath.Walk(fuzzingDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		println("Parsing: ", path)

		if err != nil {
			return err // error accessing file, TODO test
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		chain.Add(parseFile(data))

		return nil
	})
	if err != nil {
		log.Println(err) // err parsing file
	}
}

func parseFile(data []byte) []string {
	return strings.Split(string(data), " ")
}
