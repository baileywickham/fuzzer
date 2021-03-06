# Mutation

To me, mutation can happen in three places: pregereration, during tokenization, and post generation. In our example, this would be mutating the input before it is fed into the markov chain, changing the tokenizer for the markov chain, or mutating the output of the markov chain.

One of the underlying assumtions here is that we are using a markov chain however there may be better methods of generating input. While I am usually skeptical of AI, this seems like the perfect place to impliment a Neural Net, especially with the large dataset we have.

So in general, we can create interesting input by mutating the source corpus, we can change how we generate input (Markov chain, NN, Genetic Algo... etc), or we can mutate the output of our generator.

## Implimentation
The first place to impliment mutation is pregereration, or as the corpus is read off disk into the markov chain. This could be implimented as a function which is called on each file in the `walkDirectory` function. This would alter the input before it is tokenized.

The next area to mutate would be the tokenizer itself. The tokenizer is only an interface, so it can be swapped out at will in `tokenizer.go`.

The final area to impliment mutation would be on the output of the generator. This would be a function which is called on the output of the `getInput()` function.

## Example Mutations
I can think of a couple mutations that may be useful. Bit flipping and perumting tokens has already been mentioned in the spec. I think padding with normally problematic characters `\0` or white space. Using unicode characeters, especially those in other languages that look similar to ascii. You could include a table of normally useful strings, like `SHOW TABLES;` or or javascript xss injection. I'm not sure other than this...

I guess judging the effectiveness of an input will be a major challenge with fuzzing. 
