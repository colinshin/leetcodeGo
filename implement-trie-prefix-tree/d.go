package implement_trie_prefix_tree

/*
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/

type TrieNode struct {
	links []*TrieNode
	isEnd bool
}

func NewNode() *TrieNode {
	// Maximum of r links to its children, where each link corresponds to one of r character values from dataset alphabet.
	// In this article we assume that f is 26, the number of lowercase latin letters.
	const r = 26
	return &TrieNode{links: make([]*TrieNode, r)}
}

func (n *TrieNode) ContainsKey(c rune) bool {
	return n.links[c-'a'] != nil
}

func (n *TrieNode) Get(c rune) *TrieNode {
	return n.links[c-'a']
}

func (n *TrieNode) Put(c rune, node *TrieNode) {
	n.links[c-'a'] = node
}

func (n *TrieNode) SetEnd() {
	n.isEnd = true
}

func (n *TrieNode) IsEnd() bool {
	return n.isEnd
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: NewNode()}
}

/** Inserts a word into the trie.

Time complexity : O(m), where m is the key length.

Space complexity : O(m)
In the worst case newly inserted key doesn't share a prefix with the the keys already inserted in the trie.
We have to add mm new nodes, which takes us O(m) space.
*/
func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		if !node.ContainsKey(c) {
			node.Put(c, NewNode())
		}
		node = node.Get(c)
	}
	node.SetEnd()
}

/** Returns if the word is in the trie.

Time complexity : O(m) In each step of the algorithm we search for the next key character.
In the worst case the algorithm performs mm operations.

Space complexity : O(1)
*/
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.IsEnd()
}

/** Returns if there is any word in the trie that starts with the given prefix.
Time complexity : O(m)

Space complexity : O(1)O(1)
*/
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(word string) *TrieNode {
	node := t.root
	for _, c := range word {
		if node.ContainsKey(c) {
			node = node.Get(c)
		} else {
			return nil
		}
	}
	return node
}
