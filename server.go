package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type resp struct {
	Input string
}

func startServer(listenPort string) {
	http.HandleFunc("/", response)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

func response(w http.ResponseWriter, r *http.Request) {
	// Gets the next markov prediction
	markov, err := getNextMarkovPred()
	j, err := json.Marshal(resp{base64.StdEncoding.EncodeToString([]byte(markov))})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func main() {
	listenPort := flag.String("port", "8080", "Port to serve on")
	fuzzingDir := flag.String("corpus-location", ".", "location of test corpus")
	flag.Parse()

	createChain(*fuzzingDir)
	startServer(*listenPort)
}
