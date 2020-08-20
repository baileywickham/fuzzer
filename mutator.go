package main

import (
	"sort"
)

type Premutator interface {
	Mutate([]byte) []byte
}

type Postmutator interface {
	// Output of getInput
	Mutate([]string) []string
}

type nonMutator struct {
}

type sortMutator struct {
}

func (_ nonMutator) Mutate(b []byte) []byte {
	return b
}

func (_ sortMutator) Mutate(s []string) []string {
	sort.Strings(s)
	return s
}
