package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

//Point is point
type Point struct {
	Elevation int
	IsPeak    bool
	Neighbors *[]Point
}

var directions = [4][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func findNextSteps(point Point, currentPath *[]int, winner *[]int) {
	*currentPath = append(*currentPath, point.Elevation)
	for _, direction := range *point.Neighbors {
		copyCurrentPath := make([]int, len(*currentPath))
		copy(copyCurrentPath, *currentPath)
		if len(*direction.Neighbors) == 0 {
			copyCurrentPath = append(copyCurrentPath, direction.Elevation)
			if len(copyCurrentPath) > len(*winner) || (len(copyCurrentPath) == len(*winner) && getDrop(copyCurrentPath) > getDrop(*winner)) {
				*winner = copyCurrentPath
			}
		} else {
			findNextSteps(direction, &copyCurrentPath, winner)
		}
	}
}

func getDrop(path []int) int {
	if len(path) == 0 {
		return 0
	}
	return path[0] - path[len(path)-1]
}

func rebuildMap(field *[][]Point) {
	maxRow := len(*field)
	maxCol := len((*field)[0])
	for rowNum := 0; rowNum < maxRow; rowNum++ {
		for cellNum := 0; cellNum < maxCol; cellNum++ {
			isPeak := true
			for _, direction := range directions {
				neighborRow := rowNum + direction[0]
				neighborCell := cellNum + direction[1]
				if neighborRow >= 0 && neighborRow < maxRow && neighborCell >= 0 && neighborCell < maxCol {
					neighbor := (*field)[neighborRow][neighborCell]
					if (*field)[rowNum][cellNum].Elevation <= neighbor.Elevation {
						isPeak = false
					} else {
						*((*field)[rowNum][cellNum]).Neighbors = append(*((*field)[rowNum][cellNum]).Neighbors, neighbor)
					}
				}
			}
			(*field)[rowNum][cellNum].IsPeak = isPeak
		}
	}
}

//Solve is solve func
func Solve(field [][]Point) []int {
	var winner []int
	rebuildMap(&field)
	maxRow := len(field)
	maxCol := len(field[0])

	for rn := 0; rn < maxRow; rn++ {
		for cn := 0; cn < maxCol; cn++ {
			if field[rn][cn].IsPeak {
				findNextSteps(field[rn][cn], &[]int{}, &winner)
			}
		}
	}
	return winner
}

//ReadMap read map from file
func ReadMap(fileName string) [][]Point {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := 0
	var field [][]Point

	for scanner.Scan() {
		inputRow := strings.Split(scanner.Text(), " ")
		if pos > 0 {
			var row []Point
			for _, cell := range inputRow {
				if cellValue, ok := strconv.Atoi(cell); ok == nil {
					row = append(row, Point{Elevation: cellValue, Neighbors: &[]Point{}})
				}
			}
			field = append(field, row)
		}
		pos++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return field
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	winner := Solve(ReadMap("./map.txt"))
	fmt.Printf("winner: %v len: %v, drop %v", winner, len(winner), getDrop(winner))
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
