package maps

import "errors"

type D map[string]string

var (
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
)

func (d D) Search(w string) (string, error) {
	definition, ok := d[w]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d D) Add(word, definition string) error {
	d[word] = definition
	return nil

}

func (d D) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
func (d D) Delete(word string) {
	delete(d, word)
}
