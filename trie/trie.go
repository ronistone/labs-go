// Package trie is a Simple trie Implementation
package trie

type Trie struct {
	value      string
	children   map[uint8]*Trie
	isTerminal bool
}

const RootValue = "#####"

// AddWord receive one word and insert that in the trie
func (t *Trie) AddWord(word string) {
	if len(word) == 0 {
		t.isTerminal = true
		return
	}

	character := word[0]
	child, ok := t.children[character]
	if !ok {
		child = &Trie{
			value:    string(character),
			children: make(map[uint8]*Trie, 25),
		}
		t.children[character] = child
	}
	child.AddWord(word[1:])
}

// ListWordsByPrefix Get all words in the trie that match with a prefix
func (t Trie) ListWordsByPrefix(prefix string) []string {
	var results []string
	var actualValue string

	if t.value == RootValue {
		actualValue = ""
	} else {
		actualValue = t.value
	}
	if len(prefix) == 0 {
		if t.isTerminal {
			results = append(results, actualValue)
		}

		for _, child := range t.children {
			childResults := child.ListWordsByPrefix(prefix)
			for _, childResult := range childResults {
				results = append(results, actualValue+childResult)
			}
		}
	} else {
		character := prefix[0]
		child, ok := t.children[character]

		if ok {
			childResults := child.ListWordsByPrefix(prefix[1:])
			for _, childResult := range childResults {
				results = append(results, actualValue+childResult)
			}
		}
	}
	return results
}

// QueryWord check if a word is in the trie
func (t Trie) QueryWord(word string) bool {
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
			result := child.QueryWord(word[1:])
			return result
		}
		return false
	}

}

// PrintTrie gets all words in the trie
func (t Trie) PrintTrie() []string {
	if len(t.children) == 0 {
		return []string{t.value}
	}

	results := make([]string, 0)
	for _, child := range t.children {
		childResults := child.PrintTrie()
		for i, result := range childResults {
			childResults[i] = t.value + result
			results = append(results, childResults[i])
		}
	}
	return results
}

// CreateTrie Initialize a empty trie
func CreateTrie() *Trie {
	return &Trie{
		value:    RootValue,
		children: map[uint8]*Trie{},
	}
}
