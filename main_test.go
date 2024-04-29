package main

import (
	"os"
	"testing"
)

func TestGetNotOk(t *testing.T) {

	s, err := Open("")
	if err != nil {
		t.Fatal(err)
	}

	_, ok := s.Get("k0")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestSetGet(t *testing.T) {

	s, err := Open("")
	if err != nil {
		t.Fatal(err)
	}

	want := "v0"
	s.Set("k0", want)

	got, _ := s.Get("k0")
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestUpdate(t *testing.T) {

	s, err := Open("")
	if err != nil {
		t.Fatal(err)
	}

	want := "v"
	s.Set("k0", "v0")
	s.Set("k0", want)

	got, _ := s.Get("k0")
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestPersistence(t *testing.T) {

	path := t.TempDir() + "/kv.data"
	_, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	s.Set("k0", "v0")

	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}

	s2, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	got, _ := s2.Get("k0")
	if got != "v0" {
		t.Errorf("want: %v, got: %v", "v0", got)
	}
}
