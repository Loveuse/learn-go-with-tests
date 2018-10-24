package dictionary

const (
	ErrKeyNotFound          = DictionaryErr("key not in the dictionary")
	ErrExistingKey          = DictionaryErr("Already existing key")
	ErrUpdateNotExistingKey = DictionaryErr("Key does not exist - can't update it")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrExistingKey
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrUpdateNotExistingKey
	case nil:
		d[key] = newValue
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
