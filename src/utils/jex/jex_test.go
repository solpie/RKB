package jex

import (
	"testing"
	. "github.com/coryb/sorty"
)


const (
	_ = iota + 1000
	cs_fadeIn2
	cs_fadeIn3
)

func TestJsonEx_Sort(t *testing.T) {
	//topicStatus()
	s := NewSorter().ByKeys([]string{
		"+foo",
		"-bar",
	})

	data := []map[string]interface{}{
		{"foo": "abc", "bar": 890},
		{"foo": "xyz", "bar": 123},
		{"foo": "def", "bar": 456},
		{"foo": "mno", "bar": 789},
		{"foo": "def", "bar": 789},
	}

	s.Sort(data)
	println(cs_fadeIn2)
	println(cs_fadeIn3)
}