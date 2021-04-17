package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	s        string
	k        int
	expected string
}

var data = []TestCase{
	{
		s:        "abcd",
		k:        2,
		expected: "abcd",
	},
	{
		s:        "deeedbbcccbdaa",
		k:        3,
		expected: "aa",
	},
	// {
	// 	s:        "pbbcggttciiippooaais",
	// 	k:        2,
	// 	expected: "ps",
	// },
	// {
	// 	s:        "",
	// 	k:        3,
	// 	expected: "",
	// },
	// {
	// 	s:        "a",
	// 	k:        1,
	// 	expected: "",
	// },
	// {
	// 	s:        "a",
	// 	k:        2,
	// 	expected: "a",
	// },
}

var solvers = []Solver{
	RewindSolver{},
	// LinearSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.RemoveDuplicates(d.s, d.k))
		}
	}
}

func TestCreateWalkSeek(t *testing.T) {
	s := ""
	assert.Equal(t, s, String(New(s), true))
	assert.Equal(t, s, String(New(s), false))
	s = "a"
	assert.Equal(t, s, String(New(s), true))
	assert.Equal(t, s[1:], String(Seek(New(s), 1), true))
	assert.Equal(t, "a", String(Seek(New(s), len(s)-1), false))
	s = "abcd"
	assert.Equal(t, s, String(New(s), true))
	assert.Equal(t, s[1:], String(Seek(New(s), 1), true))
	assert.Equal(t, "dcba", String(Seek(New(s), len(s)-1), false))
}
