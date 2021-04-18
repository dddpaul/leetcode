package main

import (
	"math/rand"
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
	{
		s:        "pbbcggttciiippooaais",
		k:        2,
		expected: "ps",
	},
	{
		s:        "nnwssswwnvbnnnbbqhhbbbhmmmlllm",
		k:        3,
		expected: "vqm",
	},
	{
		s:        "yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy",
		k:        4,
		expected: "ybth",
	},
	{
		s:        "",
		k:        3,
		expected: "",
	},
	{
		s:        "a",
		k:        1,
		expected: "",
	},
	{
		s:        "a",
		k:        2,
		expected: "a",
	},
}

var solvers = []Solver{
	ArrayRewindSolver{},
	ListRewindSolver{},
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
	last := Seek(New(s), len(s)-1)
	assert.Equal(t, "d", string(last.Val))
	assert.Equal(t, "cba", String(Seek(last, -1), false))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func benchmark(s Solver, k int, b *testing.B) {
	d := TestCase{
		s:        randSeq(1_000_000),
		k:        k,
		expected: "",
	}
	s.RemoveDuplicates(d.s, d.k)
}

func BenchmarkArray1(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 1, b)
}

func BenchmarkArray10(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 10, b)
}

func BenchmarkArray100(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 100, b)
}

func BenchmarkList1(b *testing.B) {
	benchmark(ListRewindSolver{}, 1, b)
}

func BenchmarkList10(b *testing.B) {
	benchmark(ListRewindSolver{}, 10, b)
}

func BenchmarkList100(b *testing.B) {
	benchmark(ListRewindSolver{}, 100, b)
}
