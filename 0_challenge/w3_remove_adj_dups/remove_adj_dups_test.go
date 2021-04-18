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

var letters = []rune("abc")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func benchmark(s Solver, k int, b *testing.B) {
	d := TestCase{
		s:        randSeq(1_000),
		k:        k,
		expected: "",
	}
	for i := 0; i < b.N; i++ {
		s.RemoveDuplicates(d.s, d.k)
	}
}

func BenchmarkArray2(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 2, b)
}

func BenchmarkArray3(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 3, b)
}

func BenchmarkArray4(b *testing.B) {
	benchmark(ArrayRewindSolver{}, 4, b)
}

func BenchmarkList2(b *testing.B) {
	benchmark(ListRewindSolver{}, 2, b)
}

func BenchmarkList3(b *testing.B) {
	benchmark(ListRewindSolver{}, 3, b)
}

func BenchmarkList4(b *testing.B) {
	benchmark(ListRewindSolver{}, 4, b)
}

/*
goos: darwin
goarch: amd64
pkg: github.com/dddpaul/leetcode/0_challenge/w3_remove_adj_dups
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkArray2-16    	   23611	     51462 ns/op
BenchmarkArray3-16    	   84358	     18275 ns/op
BenchmarkArray4-16    	  181688	      5942 ns/op
BenchmarkList2-16     	   10000	    103323 ns/op
BenchmarkList3-16     	    6933	    175089 ns/op
BenchmarkList4-16     	    5958	    194401 ns/op
*/
