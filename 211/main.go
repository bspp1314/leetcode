package main

import "fmt"

type WordDictionary struct {
	isWord bool
	next [26]*WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		isWord:false,
		next:[26]*WordDictionary{},
	}
}


func (d *WordDictionary) AddWord(word string)  {
	if len(word) == 0 {
		return
	}

	for i := 0;i <len(word);i++ {
		if d.next[word[i]-'a'] == nil {
			d.next[word[i] - 'a'] = &WordDictionary{
				isWord: false,
				next:   [26]*WordDictionary{},
			}
		}

		d = d.next[word[i]-'a']
	}

	d.isWord = true
}

func (d *WordDictionary) Search(word string) bool {
	length := len(word)
	if length == 0 {
		return true
	}

	return d.match(word, 0)
}

func (d *WordDictionary) match(word string,index int)bool  {
	if index == len(word) {
		return d.isWord
	}

	if word[index] == '.' {
		for i:=0;i < 26;i++ {
			if d.next[i] == nil {
				continue
			}

			if d.next[i].match(word, index+1) {
				return true
			}
		}

		return false
	}

	if d = d.next[word[index]-'a']; d == nil {
		return false
	}
	return d.match(word, index+1)

}



func main() {
	w := Constructor()
	w.AddWord("a")
	fmt.Println(w.Search("a."))
}

