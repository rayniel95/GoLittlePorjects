package prefixtree

type ITrie interface {
	addWord(word []byte)
	count() int
	prefix(prefix []byte) int
	this(char byte) *TrieNode
}

type TrieNode struct {
	value   byte
	endWord bool
	sons    map[byte]*TrieNode
}

type Trie struct {
	startWords map[byte]*TrieNode
}

func (trieNode *TrieNode) addSon(value byte, isEndWord bool) bool {
	_, ok := (*trieNode).sons[value]
	if ok {
		return false
	}
	(*trieNode).sons[value] = &TrieNode{value, isEndWord, make(map[byte]*TrieNode)}
	return true
}

func (trieNode *TrieNode) addWord(word []byte) {
	index := 0
	node := trieNode
	for index < len(word) && (*node).sons[word[index]] != nil {
		index++
		node = (*node).sons[word[index]]
	}
	if index >= len(word) {
		(*node).endWord = true
		return
	}
	for ; index < len(word); index++ {
		node.addSon(word[index], false)
		node = (*node).sons[word[index]]
	}
	(*node).endWord = true
}

func (trie *TrieNode) count() int {
	sum := 0
	for _, son := range (*trie).sons {
		sum += son.count()
	}
	if (*trie).endWord {
		return 1 + sum
	}
	return sum
}

func (trie *TrieNode) prefix(word []byte) int {
	node, ok := (*trie).sons[word[0]]
	if !ok {
		return 0
	}
	index := 1
	for ; index < len(word); index++ {
		node, ok = (*node).sons[word[index]]
		if !ok {
			return 0
		}
	}
	return node.count()
}

func (trie *TrieNode) this(char byte) *TrieNode {
	val, ok := (*trie).sons[char]
	if !ok {
		return nil
	}
	return val
}
