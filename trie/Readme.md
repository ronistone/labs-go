
#### A Simple Trie Implementation to train golang

### Run

The script expect multiples values in multiples lines to build the trie, use EOF (ctrl+d) to stop the build.


#### Output example

````
Adding Word: banana
Adding Word: batata
Adding Word: batalha
Adding Word: bolacha
Adding Word: (EOF)
Consulting words: ba
ba isFound:  false

Listing By Prefix
batalha
batata
banana

Consulting words: b
b isFound:  false

Listing By Prefix
batata
batalha
banana
bolacha

Consulting words: c
c isFound:  false

Listing By Prefix

Consulting words: (EOF)

````