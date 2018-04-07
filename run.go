package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getNeighbors(row int, cell int, field [][]int) [][]int {
	var neighbors [][]int
	directions := [4][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	for _, direction := range directions {
		neighborRow := row + direction[0]
		neighborCell := cell + direction[1]
		if neighborRow >= 0 && neighborRow < len(field) {
			if neighborCell >= 0 && neighborCell < len(field[0]) {
				neighbors = append(neighbors, []int{neighborRow, neighborCell})
			}
		}
	}
	return neighbors
}
func fundPeaks(field [][]int) [][]int {
	var data [][]int
	for rowNum, row := range field {
		// fmt.Printf("rowNum: %v\n", rowNum)
		for cellNum := range row {
			// fmt.Printf("cellNum: %v\n", cellNum)
			neighbors := getNeighbors(rowNum, cellNum, field)
			// fmt.Printf("cell: %v(%vx%v), neighbors: %v\n", field[rowNum][cellNum], rowNum, cellNum, neighbors)
			ok := true
			for _, neighbor := range neighbors {
				if field[neighbor[0]][neighbor[1]] >= field[rowNum][cellNum] {
					ok = false
				}
			}
			if ok {
				data = append(data, []int{rowNum, cellNum})
			}

		}
	}
	return data
}
func getPossibleDirections(row int, cell int, field [][]int) [][]int {
	var directions [][]int
	neighbors := getNeighbors(row, cell, field)
	for _, neighbor := range neighbors {
		if field[neighbor[0]][neighbor[1]] < field[row][cell] {
			directions = append(directions, neighbor)
		}
	}
	return directions
}
func findNextSteps(row int, cell int, field [][]int, currentPath *[]int, results *[][]int) {
	*currentPath = append(*currentPath, field[row][cell])
	possibleDirections := getPossibleDirections(row, cell, field)
	for _, direction := range possibleDirections {
		copyCurrentPath := make([]int, len(*currentPath))
		copy(copyCurrentPath, *currentPath)
		nextMoves := getPossibleDirections(direction[0], direction[1], field)
		if len(nextMoves) == 0 {
			copyCurrentPath = append(copyCurrentPath, field[direction[0]][direction[1]])
			*results = append(*results, copyCurrentPath)
		} else {
			findNextSteps(direction[0], direction[1], field, &copyCurrentPath, results)
		}
	}
}

type byLength [][]int

func getDrop(path []int) int {
	return path[0] - path[len(path)-1]
}

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func solve(field [][]int) {
	peaks := fundPeaks(field)
	var results [][]int
	for _, peak := range peaks {
		findNextSteps(peak[0], peak[1], field, &[]int{}, &results)
	}
	if len(results) > 0 {
		sort.Sort(byLength(results))
		winner := results[len(results)-1]
		fmt.Printf("winner: %v len: %v, drop %v", winner, len(winner), getDrop(winner))
	}

}
func main() {
	file, err := os.Open("./map.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := 0
	var field [][]int

	for scanner.Scan() {
		inputRow := strings.Split(scanner.Text(), " ")
		if pos > 0 {
			var row []int
			for _, cell := range inputRow {
				if cellValue, ok := strconv.Atoi(cell); ok == nil {
					row = append(row, cellValue)
				}
			}
			field = append(field, row)
		}
		pos++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	solve(field)
}
