package main

import "fmt"

type trie struct {
	value      string
	children   map[uint8]*trie
	isTerminal bool
}

func (t *trie) addWord(word string) {
	if len(word) == 0 {
		t.isTerminal = true
		return
	}

	character := word[0]
	child, ok := t.children[character]
	if !ok {
		child = &trie{
			value:    string(character),
			children: make(map[uint8]*trie, 25),
		}
		t.children[character] = child
	}
	child.addWord(word[1:])
}

func (t trie) listWordsByPrefix(prefix string) []string {
	var results []string

	if len(prefix) == 0 {
		if t.isTerminal {
			results = append(results, t.value)
		}

		for _, child := range t.children {
			childResults := child.listWordsByPrefix(prefix)
			for _, childResult := range childResults {
				results = append(results, t.value+childResult)
			}
		}
	} else {
		character := prefix[0]
		child, ok := t.children[character]

		if ok {
			childResults := child.listWordsByPrefix(prefix[1:])
			for _, childResult := range childResults {
				results = append(results, t.value+childResult)
			}
		}
	}
	return results
}

func (t trie) queryWord(word string) bool {
	if len(word) == 0 {
		if t.isTerminal {
			return true
		} else {
			return false
		}
	} else {
		character := word[0]
		child, ok := t.children[character]

		if ok {
			result := child.queryWord(word[1:])
			return result
		}
		return false
	}

}

func (t trie) printTrie() []string {
	if len(t.children) == 0 {
		return []string{t.value}
	}

	results := make([]string, 0)
	for _, child := range t.children {
		childResults := child.printTrie()
		for i, result := range childResults {
			childResults[i] = t.value + result
			results = append(results, childResults[i])
		}
	}
	return results
}

func main() {
	data := &trie{
		children: map[uint8]*trie{},
	}
	var input string
	for {
		fmt.Printf("Adding Word: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		data.addWord(input)
	}

	for {
		fmt.Printf("\nConsulting words: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		resultQuery := data.queryWord(input)
		fmt.Println(input, "isFound: ", resultQuery)

		fmt.Println("\nListing By Prefix")
		results := data.listWordsByPrefix(input)
		for _, word := range results {
			fmt.Println(word)
		}
	}

	fmt.Println("Printing all words:")
	for _, word := range data.printTrie() {
		fmt.Println(word)
	}
}
