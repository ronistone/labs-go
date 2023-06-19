package main

import (
	"fmt"
	"github.com/ronistone/labs-go/trie"
)

func main() {
	data := trie.CreateTrie()
	var input string
	for {
		fmt.Printf("Adding Word: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		data.AddWord(input)
	}

	for {
		fmt.Printf("\nConsulting words: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		resultQuery := data.QueryWord(input)
		fmt.Println(input, "isFound: ", resultQuery)

		fmt.Println("\nListing By Prefix")
		results := data.ListWordsByPrefix(input)
		for _, word := range results {
			fmt.Println(word)
		}
	}

	fmt.Println("Printing all words:")
	for _, word := range data.PrintTrie() {
		fmt.Println(word)
	}
}
