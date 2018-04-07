package main

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestReadMap(t *testing.T) {
	g := Goblin(t)
	g.Describe("read file", func() {

		tmpfile, _ := ioutil.TempFile(os.TempDir(), "prefix")
		tmpfile.Write([]byte("4 4\n4 8 7 3\n2 5 9 3\n6 3 2 5\n4 4 1 6"))
		defer os.Remove(tmpfile.Name())

		field := ReadMap(tmpfile.Name())
		testField := [][]int{
			{4, 8, 7, 3},
			{2, 5, 9, 3},
			{6, 3, 2, 5},
			{4, 4, 1, 6},
		}

		g.It("Should be equal ", func() {
			g.Assert(field).Equal(testField)
		})
	})
}

func TestSolve(t *testing.T) {
	g := Goblin(t)
	g.Describe("solve correctly", func() {
		testField := [][]int{
			{4, 8, 7, 3},
			{2, 5, 9, 3},
			{6, 3, 2, 5},
			{4, 4, 1, 6},
		}
		testAnswer := []int{8, 5, 3, 2, 1}

		g.It("Should solve correctly ", func() {
			g.Assert(Solve(testField)).Equal(testAnswer)
		})
	})
}

func TestFindPeaks(t *testing.T) {
	g := Goblin(t)
	g.Describe("find peaks ", func() {
		testField := [][]int{
			{4, 8, 7, 3},
			{2, 5, 9, 3},
			{6, 3, 2, 5},
			{4, 4, 1, 6},
		}
		testAnswer := [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 3}}

		g.It("Should find all peaks ", func() {
			g.Assert(FindPeaks(testField)).Equal(testAnswer)
		})
	})
}

func TestGetNeighbors(t *testing.T) {
	g := Goblin(t)
	g.Describe("find GetNeighbors  ", func() {
		testField := [][]int{
			{4, 8, 7, 3},
			{2, 5, 9, 3},
			{6, 3, 2, 5},
			{4, 4, 1, 6},
		}
		g.It("Should find neighbors ", func() {
			g.Assert(GetNeighbors(-1, -1, testField)).Equal([][]int(nil))
			g.Assert(GetNeighbors(0, 0, testField)).Equal([][]int{[]int{1, 0}, []int{0, 1}})
			g.Assert(GetNeighbors(0, 1, testField)).Equal([][]int{[]int{0, 0}, []int{1, 1}, []int{0, 2}})
			g.Assert(GetNeighbors(0, 2, testField)).Equal([][]int{[]int{0, 1}, []int{1, 2}, []int{0, 3}})
			g.Assert(GetNeighbors(0, 3, testField)).Equal([][]int{[]int{0, 2}, []int{1, 3}})
			g.Assert(GetNeighbors(1, 0, testField)).Equal([][]int{[]int{0, 0}, []int{2, 0}, []int{1, 1}})
			g.Assert(GetNeighbors(2, 2, testField)).Equal([][]int{[]int{1, 2}, []int{2, 1}, []int{3, 2}, []int{2, 3}})
		})
	})
}
