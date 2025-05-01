package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"key": "value"}
	t.Run("known key", func(t *testing.T) {

		got, _ := d.Search("key")
		want := "value"
		assertStrings(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		_, err := d.Search("")
		if err == nil {
			t.Fatal("expected an error", err)
		}
		assertStrings(t, err.Error(), ErrKeyNotFound.Error())
	})

}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	t.Run("add key", func(t *testing.T) {

		d.Add("key", "value")
		want := "value"
		got, err := d.Search("key")
		if err != nil {
			t.Fatal("not found added key", err)
		}
		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
