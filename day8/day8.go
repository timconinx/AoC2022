package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var grid [][]int = make([][]int, 0)
var x, y int

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := strings.Split(reader.Text(), "")
		if x == 0 {
			x = len(text)
		}
		grid = append(grid, make([]int, x))
		for i := 0; i < x; i++ {
			number, _ := strconv.Atoi(text[i])
			grid[y][i] = number
		}
		y++
	}
	visible := 2*x + 2*(y-2)
	for i := 1; i < x-1; i++ {
		for j := 1; j < y-1; j++ {
			if isVisible(i, j) {
				visible++
			}
		}
	}
	maxssc := 1
	for i := 1; i < x-1; i++ {
		for j := 1; j < y-1; j++ {
			var ssc = scenicScore(i, j)
			if ssc > maxssc {
				maxssc = ssc
			}
		}
	}
	println("total trees visible is " + strconv.Itoa(visible))
	println("max scenic score is " + strconv.Itoa(maxssc))
}

func scenicScore(a, b int) int {
	var scoreFromTop, scoreFromBottom, scoreFromLeft, scoreFromRight int
	for i := a - 1; i > -1; i-- {
		// top
		scoreFromTop++
		if grid[i][b] >= grid[a][b] {
			break
		}
	}
	for i := a + 1; i < x; i++ {
		// bottom
		scoreFromBottom++
		if grid[i][b] >= grid[a][b] {
			break
		}
	}
	for j := b - 1; j > -1; j-- {
		// left
		scoreFromLeft++
		if grid[a][j] >= grid[a][b] {
			break
		}
	}
	for j := b + 1; j < y; j++ {
		// right
		scoreFromRight++
		if grid[a][j] >= grid[a][b] {
			break
		}
	}
	return scoreFromTop * scoreFromBottom * scoreFromLeft * scoreFromRight
}

func isVisible(a, b int) bool {
	var visibleFromTop, visibleFromBottom, visibleFromLeft, visibleFromRight bool = true, true, true, true
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i < a && j == b {
				// top
				visibleFromTop = visibleFromTop && (grid[i][j] < grid[a][b])
			}
			if i > a && j == b {
				// bottom
				visibleFromBottom = visibleFromBottom && (grid[i][j] < grid[a][b])
			}
			if i == a && j < b {
				// left
				visibleFromLeft = visibleFromLeft && (grid[i][j] < grid[a][b])
			}
			if i == a && j > b {
				// right
				visibleFromRight = visibleFromRight && (grid[i][j] < grid[a][b])
			}
		}
	}
	return visibleFromTop || visibleFromBottom || visibleFromLeft || visibleFromRight
}
