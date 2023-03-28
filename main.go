package main

import (
	"encoding/gob"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// single
	gob.Register(Tiny{})
	gob.Register(Medium{})
	gob.Register(Huge{})
	// slices
	gob.Register([]Tiny{})
	gob.Register([]Medium{})
	gob.Register([]Huge{})
	// slices of pointers
	gob.Register([]*Tiny{})
	gob.Register([]*Medium{})
	gob.Register([]*Huge{})

	os.Exit(m.Run())
}
