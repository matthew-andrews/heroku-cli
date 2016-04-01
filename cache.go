package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type cacheFetchFn func(*interface{}) error

func cacheFetch(key string, o interface{}, expiresIn time.Duration, fetch cacheFetchFn) error {
	filename := filepath.Join(AppDir(), "cache", key)

	read := func(o interface{}) (cached bool, err error) {
		fi, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		if time.Since(fi.ModTime()) > expiresIn {
			return false, nil
		}
		file, err := os.Open(filename)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		if err := json.NewDecoder(file).Decode(&o); err != nil {
			return false, err
		}
		return true, nil
	}

	save := func(o interface{}) error {
		data, err := json.MarshalIndent(o, "", "  ")
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return err
		}
		return ioutil.WriteFile(filename, data, 0644)
	}

	cached, err := read(&o)
	if err != nil {
		return err
	}
	if !cached {
		p := o
		if err := fetch(&p); err != nil {
			return err
		}
		if err := save(p); err != nil {
			return err
		}
		if _, err := read(&o); err != nil {
			return err
		}
	}
	return nil
}
