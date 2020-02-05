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

type Node struct {
	links    []*Node
	IsEnd    bool
	linkNums int
}

func NewNode() *Node {
	// Maximum of r --l--inks to its children, where each link corresponds to one of r character values from dataset alphabet.
	// In this article we assume that r is 26, the number of lowercase latin letters.
	const r = 26
	return &Node{links: make([]*Node, r)}
}

func (n *Node) Has(c byte) bool {
	return n.links[c-'a'] != nil
}

func (n *Node) Get(c byte) *Node {
	return n.links[c-'a']
}

func (n *Node) Put(c byte, node *Node) {
	if !n.Has(c) {
		n.linkNums++
	}
	n.links[c-'a'] = node
}

type Trie struct {
	root *Node
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
	p := t.root
	for i := 0; i < len(word); i++ {
		if !p.Has(word[i]) {
			p.Put(word[i], NewNode())
		}
		p = p.Get(word[i])
	}
	p.IsEnd = true
}

/** Returns if the word is in the trie.

Time complexity : O(m) In each step of the algorithm we search for the next key character.
In the worst case the algorithm performs mm operations.

Space complexity : O(1)
*/
func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix.
Time complexity : O(m)

Space complexity : O(1)
*/
func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}

func (t *Trie) search(s string) *Node {
	node := t.root
	for i := 0; i < len(s); i++ {
		if node.Has(s[i]) {
			node = node.Get(s[i])
		} else {
			return nil
		}
	}
	return node
}

/*
Time complexity : O(m) In each step of the algorithm we search for the next key character.
In the worst case the algorithm performs m operations.

Space complexity : O(1)
*/
func (t *Trie) SearchLongestPrefixOf(word string) string {
	k := 0
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if !node.Has(ch) || node.linkNums == 1 || node.IsEnd {
			return word[:k]
		}
		k++
		node = node.Get(ch)
	}
	return word[:k]
}
