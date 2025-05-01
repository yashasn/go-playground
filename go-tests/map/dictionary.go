package main

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("Key not found")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
