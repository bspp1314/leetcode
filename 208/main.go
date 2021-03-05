package main

import "fmt"

type Trie struct {
	isWord bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isWord:false,
		next:[26]*Trie{},
	}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	if len(word) == 0 {
		return
	}

	for i := 0;i <len(word);i++ {
		if t.next[word[i]-'a'] == nil {
			t.next[word[i] - 'a'] = &Trie{
				isWord: false,
				next:   [26]*Trie{},
			}
		}

		t = t.next[word[i]-'a']
	}

	t.isWord = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	return t.match(word,0,false)
}

func (t *Trie) match(word string,index int,startWith bool) bool  {
	if index == len(word)  {
		return t.isWord || startWith
	}



	if t = t.next[word[index]-'a']; t == nil {
		return false
	}
	return t.match(word, index+1,startWith)
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.match(prefix,0,true)
}

func main() {
	 w := Constructor()
	 w.Insert("startWith")

	 fmt.Println(w.StartsWith("start"))
	 fmt.Println(w.Search("start"))
	fmt.Println(w.Search("startWith"))

}


