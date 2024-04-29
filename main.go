package main

import (
	"encoding/gob"
	"errors"
	"io"
	"io/fs"
	"log"
	"log/slog"
	"os"
)

type Store struct {
	data map[string]string
	path string
}

func Open(filepath string) (*Store, error) {

	s := Store{
		data: make(map[string]string),
		path: filepath,
	}

	f, err := os.Open(filepath)
	// If file does not exist, assume inmem usage and return created store.
	if errors.Is(err, fs.ErrNotExist) {
		return &s, nil
	}
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil && err != io.EOF {
		log.Fatal("encode error:", err)
	}

	return &s, nil
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s Store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	if !ok {
		return "", false
	}
	return v, true
}

func (s Store) Save() error {

	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = gob.NewEncoder(f).Encode(s.data)
	if err != nil {
		return err
	}

	slog.Info("data persisted to file", "path", s.path)

	return nil
}
