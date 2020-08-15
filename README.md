# fuzzbuzz
Interview question for FuzzBuzz

Built go 1.14 on Linux Mint

## Design
The program is roughly split into three compontents:
- server.go
- markov.go
- tokenizer.go

The goal of this was to provide some sense of modularity. In the future, you may want to use a Neural Net instead of a markov chain, or you may want to use a different tokenizer for MP3 files vs http files, and the modular design allows for easier transitions between these modes.
## FAQ

### Why are you not using goroutines?
While the parsing would be a good place to impliment goroutines, I felt the like the complexity was too much for this short of a project. Especially because parsing all of mozilla is still < 1 second.

### Why does your project not conform to the design spec given?
My program does not use the `--corpus-location` paramater, instead all arguments are assumed to be locations by default. This is because I wanted to handle having one server with multiple markov chains, and it was simplest to use all arguments as locations.
