package trie

import "testing"

func recursiveAssert(t *testing.T, head string, word string, trie *Trie) {
	if len(word) == 0 {
		return
	}
	if len(trie.children) != 1 {
		t.Errorf("Must Have Only one Child")
	}

	if trie.value != head {
		t.Errorf("Value of trie node is (%s) different that wanted (%s)!", trie.value, string(word[0]))
	}

	if child, ok := trie.children[word[0]]; ok {
		recursiveAssert(t, string(word[0]), word[1:], child)
		return
	} else {
		t.Errorf("The word inserted must have %s", word)
	}
}

func TestTrie_AddWord(t *testing.T) {
	trie := CreateTrie()

	trie.AddWord("batata")

	recursiveAssert(t, RootValue, "batata", trie)
}
