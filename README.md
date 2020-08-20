# Fuzzer
Built go 1.14 on Linux Mint

See the **WAV** branch for half baked WAV tokenizing support.

See the **mutation** branch for the mutation implimentation.

See the mutation branch for an implimentation of mutation.

## Use

### Generation
The easiest way to use this program is through docker. The `dock.sh` runs the program against the mozilla application on port 8080 by default.
```bash
./dock.sh
curl localhost:8080/fuzzing-corpus/xml/mozilla
```

You can also run on your local machine. The program requires [gomarkov](https://github.com/mb-14/gomarkov), which should be automaticly downloaded using `go run .`. The default port is `:8080`.
```golang
go run . -port=8080 corpus-directory1 corpus-directory2
```


One of the features of my implimentation is allowing for multiple markov chains in the same server. By default, they are given as a list of arguments, with each entry coresponding to a directory which will be fed into a markov chain. They are then automaticly served on their directory name. For example:
```golang
go run . -port=8080 corpus-directory1 corpus-directory2
curl localhost:8080/corpus-directory1
curl localhost:8080/corpus-directory2
```
This code will create two markov chains, one based on corpus-directory1, and the second based on corpus-directory2. These are indepent markov chains.

There is also an option in the code to change the directory on which these chains are served.

### Live Updates
Submit a json post request to api endpoint coresponding to the chain you want to modify to add the new input to the chain and save the data for later use.

The post request should follow the same form as the input:
```json
{"Input": "base64..."}
```

See `post.sh` for an example of how to post json using curl. `post.sh` contains generated input which can be used to check the "live updates" functionality.

## Design
The program is roughly split into three compontents:
- server.go
- markov.go
- tokenizer.go

The goal of this was to provide some sense of modularity. In the future, you may want to use a Neural Net instead of a markov chain, or you may want to use a different tokenizer for MP3 files vs http files, and the modular design allows for easier transitions between these modes.

## FAQ

### Why are you not using goroutines?
While the parsing would be a good place to impliment goroutines, I felt the like the complexity was too much for this short of a project. Especially because parsing all of mozilla is still < 1 second.

### Can I change the location a chain is served on? What about using a different tokenizer?
Well you can change both of them, but that involves changing the source code. Ideally, these would both be command line options, but that would take longer to impliment. Ideally, you could extend this server as a package, which would allow for easy modification.
