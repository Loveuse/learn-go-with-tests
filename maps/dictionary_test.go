package dictionary

import (
	"testing"
)

const (
	key   = "key"
	value = "this is a test"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{key: value}
	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search(key)

		assertString(t, got, value, key)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("untest")

		assertError(t, err, ErrKeyNotFound, "untest")
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a non existing key", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add(key, value)

		assertError(t, err, nil, value)
		assertKey(t, dict, key, value)
	})

	t.Run("add an existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrExistingKey, key)
		assertKey(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		newValue := "this is a new test"
		err := dictionary.Update(key, newValue)

		assertError(t, err, nil, key)
		assertKey(t, dictionary, key, newValue)
	})

	t.Run("update a non existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		newValue := "this is a new test"
		err := dictionary.Update(key, newValue)

		assertError(t, err, ErrUpdateNotExistingKey, key)
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete an existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		dictionary.Delete(key)

		_, errSearch := dictionary.Search(key)

		if errSearch != ErrKeyNotFound {
			t.Errorf("Key can't be deleted")
		}
	})
}

func assertString(t *testing.T, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s given %s", got, want, given)
	}
}

func assertError(t *testing.T, err, wantedError error, given string) {
	t.Helper()
	if err != wantedError {
		t.Errorf("got %s want %s", err, wantedError)
	}

}

func assertKey(t *testing.T, dictionary Dictionary, key, want string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Errorf("should find %s", key)
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
