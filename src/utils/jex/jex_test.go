package jex

import (
	"testing"
	. "github.com/coryb/sorty"
)

func TestJsonEx_Sort(t *testing.T) {
	s := NewSorter().ByKeys([]string{
		"+foo",
		"-bar",
	})

	data := []JsonEx2{
		{"foo": "abc", "bar": 890},
		{"foo": "xyz", "bar": 123},
		{"foo": "def", "bar": 456},
		{"foo": "mno", "bar": 789},
		{"foo": "def", "bar": 789},
	}

	s.Sort(data)

}