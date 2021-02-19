package main

import (
	"fmt"

	"github.com/Chung-god/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first"}
	baseWord := "test"
	dictionary.Add(baseWord, "first")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
