package main

import (
	"os"
	"testing"
)

func TestPathUnwritable(t *testing.T) {

	s, err := Open("bogus/data.store")
	if err != nil {
		t.Fatal(err)
	}

	err = s.Save()
	if err == nil {
		t.Fatal("no error for unwritable path")
	}
}

func TestInvalidData(t *testing.T) {
	_, err := Open("testdata/invalid.store")
	if err == nil {
		t.Fatal("no error for invalid data")
	}
}

func TestWrongPermissions(t *testing.T) {

	path := t.TempDir() + "/kv.data"

	_, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chmod(path, 0o000)
	if err != nil {
		t.Fatal(err)
	}

	_, err = Open(path)
	if err == nil {
		t.Fatal("no error")
	}

}

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

func TestValueFound(t *testing.T) {

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

	got, ok := s.Get("k0")
	if !ok {
		t.Fatal("key not found")
	}
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestPersistence(t *testing.T) {

	path := t.TempDir() + "/kv.data"

	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")

	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}

	s2, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	got, _ := s2.Get("A")
	if got != "1" {
		t.Errorf("want: %v, got: %v", "1", got)
	}

	got, _ = s2.Get("B")
	if got != "2" {
		t.Errorf("want: %v, got: %v", "2", got)
	}

	got, _ = s2.Get("C")
	if got != "3" {
		t.Errorf("want: %v, got: %v", "3", got)
	}
}
