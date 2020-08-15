package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

// Only used as the json response for the API
type mkResp struct {
	Input string
}

func (h *chainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mk, err := h.chain.Generate(h.getSeed())
		log.Println(mk)
		// Converts the markov generated text to base64 then to json
		j, err := json.Marshal(mkResp{base64.StdEncoding.EncodeToString([]byte(mk))})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	case "POST":
		println("hello")
	}
}

func startServers(listenPort string, fuzzingCorpi []string) {
	for _, corpi := range fuzzingCorpi {
		// Use tokenizeBySpaces by default
		// Other tokenizers are in tokenizer.go
		createEndpoint(corpi, corpi, tokenizeBySpaces{})
	}

	log.Println("Listening on :" + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

// Creates a markov chain and a coresponding API endpoink
func createEndpoint(corpi, url string, t Tokenizer) {
	h := createChain(corpi, t)
	http.Handle("/"+url, h)
	log.Println("Serving ", corpi, "on /"+corpi)
}

func main() {
	listenPort := flag.String("port", "8080", "Port to serve on")
	flag.Parse()

	fuzzingCorpi := flag.Args()
	startServers(*listenPort, fuzzingCorpi)
}
