package main

import "fmt"

type Trie struct {
	nexts [128]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, b := range word {
		if cur.nexts[b] == nil {
			cur.nexts[b] = &Trie{}
		}
		cur = cur.nexts[b]
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, b := range word {
		if this.nexts[b] != nil {
			this = this.nexts[b]
		} else {
			return false
		}
	}
	if this.isEnd {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, b := range prefix {
		if this.nexts[b] != nil {
			this = this.nexts[b]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("abc")
	param_2 := obj.Search("abc")
	param_3 := obj.StartsWith("a")
	fmt.Println(param_2, param_3)
}
