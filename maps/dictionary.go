package maps

var (
	ErrUnknownWord      = DictionaryErr("unknown word")
	ErrRedefineWord     = DictionaryErr("word redefinition")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if v, ok := d[word]; ok {
		return v, nil
	}
	return "", ErrUnknownWord
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnknownWord:
		d[word] = definition
	case nil:
		return ErrRedefineWord
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefiniton string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnknownWord:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefiniton
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnknownWord:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
