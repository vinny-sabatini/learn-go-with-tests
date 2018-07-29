package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrWordNotFound     = DictionaryErr("could not find word you were looking for")
	ErrWordExists       = DictionaryErr("word is already defined")
	ErrWordDoesNotExist = DictionaryErr("can not update a word that does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
		return nil
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}
