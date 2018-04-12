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
		testField := [][]Point{
			{Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 8, Neighbors: &[]Point{}}, Point{Elevation: 7, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}},
			{Point{Elevation: 2, Neighbors: &[]Point{}}, Point{Elevation: 5, Neighbors: &[]Point{}}, Point{Elevation: 9, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}},
			{Point{Elevation: 6, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}, Point{Elevation: 2, Neighbors: &[]Point{}}, Point{Elevation: 5, Neighbors: &[]Point{}}},
			{Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 1, Neighbors: &[]Point{}}, Point{Elevation: 6, Neighbors: &[]Point{}}},
		}

		g.It("Should be equal ", func() {
			g.Assert(field).Equal(testField)
		})
	})
}

func TestSolve(t *testing.T) {
	g := Goblin(t)
	g.Describe("solve correctly", func() {
		testField := [][]Point{
			{Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 8, Neighbors: &[]Point{}}, Point{Elevation: 7, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}},
			{Point{Elevation: 2, Neighbors: &[]Point{}}, Point{Elevation: 5, Neighbors: &[]Point{}}, Point{Elevation: 9, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}},
			{Point{Elevation: 6, Neighbors: &[]Point{}}, Point{Elevation: 3, Neighbors: &[]Point{}}, Point{Elevation: 2, Neighbors: &[]Point{}}, Point{Elevation: 5, Neighbors: &[]Point{}}},
			{Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 4, Neighbors: &[]Point{}}, Point{Elevation: 1, Neighbors: &[]Point{}}, Point{Elevation: 6, Neighbors: &[]Point{}}},
		}
		testAnswer := []int{9, 5, 3, 2, 1}

		g.It("Should solve correctly ", func() {
			g.Assert(Solve(testField)).Equal(testAnswer)
		})
	})
}
