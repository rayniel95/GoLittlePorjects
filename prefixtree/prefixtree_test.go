package prefixtree

import (
	"testing"
)

func Test_main(t *testing.T) {
	trie := initTrie()

	trie.addWord([]byte("abaco"))
	trie.addWord([]byte("abril"))
	trie.addWord([]byte("abrir"))
	trie.addWord([]byte("hormiga"))
	trie.addWord([]byte("azulejo"))
	trie.addWord([]byte("humo"))
	trie.addWord([]byte("horno"))
	trie.addWord([]byte("hornos"))

	want := 2
	got := trie.prefix([]byte("abr"))

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
