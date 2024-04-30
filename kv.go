package kv

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"log/slog"
	"os"
)

type store struct {
	data map[string]string
	path string
}

func Open(filepath string) (*store, error) {

	s := store{
		data: make(map[string]string),
		path: filepath,
	}

	f, err := os.Open(filepath)
	// If file does not exist, assume inmem usage and return created store.
	if errors.Is(err, fs.ErrNotExist) {
		return &s, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *store) Set(key, value string) {
	s.data[key] = value
}

func (s store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	if !ok {
		return "", false
	}
	return v, true
}

func (s store) Save() error {

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
