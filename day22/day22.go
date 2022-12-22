package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	NOTILE = " "
	EMPTY  = "."
	WALL   = "#"
	RIGHT  = 0
	DOWN   = 1
	LEFT   = 2
	UP     = 3
)

var grid [][]string = make([][]string, 0)
var y, x int
var currx, curry int
var currdirection int = RIGHT

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for {
		reader.Scan()
		text := reader.Text()
		if text == "" {
			break
		}
		gridline := strings.Split(text, "")
		grid = append(grid, gridline)
		if y == 0 {
			x = len(gridline)
		}
		if y != 0 && len(gridline) > x {
			padding := make([]string, len(gridline)-x)
			for i := 0; i < len(padding); i++ {
				padding[i] = " "
			}
			for i := 0; i < y; i++ {
				grid[i] = append(grid[i], padding...)
			}
			x = len(gridline)
		}
		if x > len(gridline) {
			padding := make([]string, x-len(gridline))
			for i := 0; i < len(padding); i++ {
				padding[i] = " "
			}
			grid[y] = append(grid[y], padding...)
		}
		y++
	}
	reader.Scan()
	directions := reader.Text()
	for {
		distancetxt := regexp.MustCompile(`^(\d+)`).FindStringSubmatch(directions)[1]
		directions = strings.TrimLeft(directions, distancetxt)
		distance, _ := strconv.Atoi(distancetxt)
		walk(distance)
		if directions == "" {
			break
		}
		turning := regexp.MustCompile(`^(\w)`).FindStringSubmatch(directions)[1]
		directions = directions[1:]
		currdirection = turn(turning)
	}
	password := (curry+1)*1000 + (currx+1)*4 + currdirection
	println("the password is " + strconv.Itoa(password))
}

func walk(d int) {
	var tile string
	switch currdirection {
	case RIGHT:
		for i := 0; i < d; i++ {
			destx := currx + 1
			if destx == x || grid[curry][destx] == NOTILE {
				destx, tile = firstTileLeft()
				if tile == WALL {
					break
				}
			}
			switch grid[curry][destx] {
			case EMPTY:
				currx = destx
			case WALL:
				break
			case NOTILE:
				panic("can't happen")
			}
		}
	case LEFT:
		for i := 0; i < d; i++ {
			destx := currx - 1
			if destx < 0 || grid[curry][destx] == NOTILE {
				destx, tile = firstTileRight()
				if tile == WALL {
					break
				}
			}
			switch grid[curry][destx] {
			case EMPTY:
				currx = destx
			case WALL:
				break
			case NOTILE:
				panic("can't happen")
			}
		}
	case DOWN:
		for i := 0; i < d; i++ {
			desty := curry + 1
			if desty == y || grid[desty][currx] == NOTILE {
				desty, tile = firstTileDown()
				if tile == WALL {
					break
				}
			}
			switch grid[desty][currx] {
			case EMPTY:
				curry = desty
			case WALL:
				break
			case NOTILE:
				panic("can't happen")
			}
		}
	case UP:
		for i := 0; i < d; i++ {
			desty := curry - 1
			if desty < 0 || grid[desty][currx] == NOTILE {
				desty, tile = firstTileUp()
				if tile == WALL {
					break
				}
			}
			switch grid[desty][currx] {
			case EMPTY:
				curry = desty
			case WALL:
				break
			case NOTILE:
				panic("can't happen")
			}
		}
	}
}

func turn(turning string) int {
	switch turning {
	case "R":
		switch currdirection {
		case RIGHT:
			return DOWN
		case DOWN:
			return LEFT
		case LEFT:
			return UP
		case UP:
			return RIGHT
		default:
			panic("current direction unknown")
		}
	case "L":
		switch currdirection {
		case RIGHT:
			return UP
		case DOWN:
			return RIGHT
		case LEFT:
			return DOWN
		case UP:
			return LEFT
		default:
			panic("current direction unknown")
		}
	default:
		panic("Turning unknown " + turning)
	}
}

func firstTileLeft() (int, string) {
	for i := 0; i < x; i++ {
		if grid[curry][i] != NOTILE {
			return i, grid[curry][i]
		}
	}
	panic("no first tile left")
}

func firstTileRight() (int, string) {
	for i := x - 1; i >= 0; i-- {
		if grid[curry][i] != NOTILE {
			return i, grid[curry][i]
		}
	}
	panic("no first tile right")
}

func firstTileDown() (int, string) {
	for i := 0; i < y; i++ {
		if grid[i][currx] != NOTILE {
			return i, grid[i][currx]
		}
	}
	panic("no first tile up")
}
func firstTileUp() (int, string) {
	for i := y - 1; i >= 0; i-- {
		if grid[i][currx] != NOTILE {
			return i, grid[i][currx]
		}
	}
	panic("no first tile up")
}
