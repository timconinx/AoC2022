package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type block struct {
	x int
	y int
	z int
}

type slice struct {
	min int
	max int
}

func (b block) exposedSides() int {
	totalX := 2
	for _, other := range append(blocksOnX[b.x+1], blocksOnX[b.x-1]...) {
		if b.y == other.y && b.z == other.z {
			totalX--
			if totalX == 0 {
				break
			}
		}
	}
	totalY := 2
	for _, other := range append(blocksOnY[b.y+1], blocksOnY[b.y-1]...) {
		if b.x == other.x && b.z == other.z {
			totalY--
			if totalY == 0 {
				break
			}
		}
	}
	totalZ := 2
	for _, other := range append(blocksOnZ[b.z+1], blocksOnZ[b.z-1]...) {
		if b.x == other.x && b.y == other.y {
			totalZ--
			if totalZ == 0 {
				break
			}
		}
	}
	return totalX + totalY + totalZ
}

var blocksOnX map[int][]block = make(map[int][]block)
var blocksOnY map[int][]block = make(map[int][]block)
var blocksOnZ map[int][]block = make(map[int][]block)
var XYSlices map[int]map[int]slice = make(map[int]map[int]slice)

func main() {
	allBlocks := []block{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		triplet := strings.Split(reader.Text(), ",")
		x, _ := strconv.Atoi(triplet[0])
		y, _ := strconv.Atoi(triplet[1])
		z, _ := strconv.Atoi(triplet[2])
		b := block{x: x, y: y, z: z}
		blocksOnX[x] = append(blocksOnX[x], b)
		blocksOnY[y] = append(blocksOnY[y], b)
		blocksOnZ[z] = append(blocksOnZ[z], b)
		allBlocks = append(allBlocks, b)
	}
	totalSides := 0
	for _, b := range allBlocks {
		totalSides += b.exposedSides()
	}
	println("total exposed sides is " + strconv.Itoa(totalSides))
}
