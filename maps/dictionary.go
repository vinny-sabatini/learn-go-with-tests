package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find word you were looking for")
	ErrWordExists = DictionaryErr("word is already defined")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}