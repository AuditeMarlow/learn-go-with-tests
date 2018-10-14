package maps

import "errors"

type Dictionary map[string]string

func Search(d Dictionary, word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("Could not find the word you were looking for.")
	}

	return definition, nil
}
