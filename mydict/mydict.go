package mydict

import (
	"errors"
	"fmt"
)

//just like alias

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errIsExist    = errors.New("Word existed")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

//type can has method

//Search word in dictionary
func (d Dictionary) Search(word string) (string, error) {
	key, exists := d[word]
	if exists {
		return key, nil
	}
	return "", errNotFound
}

//Add word in Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err != nil {
		d[word] = def
		return nil
	}
	fmt.Println(err)
	return errIsExist
}

//Update word definition
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err != nil {
		return errCantUpdate
	}
	d[word] = def
	return nil
}

//Delete word in dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errCantDelete
	}
	delete(d, word)
	return nil
}
