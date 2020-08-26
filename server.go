package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	r "github.com/baileywickham/runner"
)

// Only used as the json response for the API
type mkResp struct {
	Input string
}

func (h *chainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokens, err := h.getMutatedInput()
		// Print tokens in "human readable" form
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
		// handle "Live Updates"
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
		var url string
		if strings.HasPrefix(corpi, "./") {
			// fix awkward url when using ./directory
			url = corpi[2:]
		} else {
			// Use corpi name as default path
			url = corpi
		}
		// Use tokenizeBySpaces by default
		// Other tokenizers are in tokenizer.go
		createEndpoint(corpi, url, tokenizeBySpaces{}, nonMutator{}, sortMutator{})
	}

	log.Println("Listening on :" + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}

// Creates a markov chain and a coresponding API endpoink
func createEndpoint(corpi, url string, t Tokenizer, pre Premutator, post Postmutator) {
	h := createChain(corpi, t, pre, post)
	http.Handle("/"+url, h)
	log.Println("Serving ", corpi, "on /"+url)
}

func main() {
	//listenPort := flag.String("port", "8080", "Port to serve on")
	//flag.Parse()

	// Take each trailing input as an input corpus
	//fuzzingCorpi := flag.Args()
	//startServers(*listenPort, fuzzingCorpi)
	s := r.NewShell()
	s.Add_command(r.Command{
		Cmd:      "start",
		Callback: startServers,
		Helptext: "Start a list of servers"})
	s.Start()
}
