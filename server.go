package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
)

// Only used as the json response for the API
type mkResp struct {
	Input string
}

func (h *chainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokens, err := h.getInput()
		// Print tokens in human readable form
		log.Println(tokens)
		// Converts the markov generated text to base64 then to json
		j, err := json.Marshal(mkResp{base64.StdEncoding.EncodeToString([]byte(strings.Join(tokens, "")))})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	case "POST":
		var j mkResp
		err := json.NewDecoder(r.Body).Decode(&j)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := base64.StdEncoding.DecodeString(j.Input)
		err = h.writeNewEntry(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func startServers(listenPort string, fuzzingCorpi []string) {
	for _, corpi := range fuzzingCorpi {
		// Use tokenizeBySpaces by default
		// Use corpi name as default path
		// Other tokenizers are in tokenizer.go
		if strings.HasPrefix(corpi, "./") {
			// fix awkward url when using ./directory
			createEndpoint(corpi, corpi[2:], tokenizeBySpaces{})
		} else {
			createEndpoint(corpi, corpi, tokenizeBySpaces{})
		}
	}

	log.Println("Listening on :" + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

// Creates a markov chain and a coresponding API endpoink
func createEndpoint(corpi, url string, t Tokenizer) {
	h := createChain(corpi, t)
	http.Handle("/"+url, h)
	log.Println("Serving ", corpi, "on /"+url)
}

func main() {
	listenPort := flag.String("port", "8080", "Port to serve on")
	flag.Parse()

	fuzzingCorpi := flag.Args()
	createEndpoint("./fuzzing-corpus/wav/mozilla/", "wav", tokenizeWav{})
	startServers(*listenPort, fuzzingCorpi)
}
