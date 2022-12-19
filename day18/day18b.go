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

func (s slice) adapt(v int) slice {
	if v < s.min || s.min == 0 {
		s.min = v
	}
	if v > s.max {
		s.max = v
	}
	return s
}

func (b block) exposedSides() int {
	totalX := 2
	totalY := 2
	totalZ := 2
	xyslice := XYSlices[b.x][b.y]
	yzslice := YZSlices[b.y][b.z]
	xzslice := XZSlices[b.x][b.z]
	if b.x < yzslice.max {
		if b.x+1 > XYSlices[b.x+1][b.y].min && b.x+1 < XYSlices[b.x+1][b.y].max {
			if b.x+1 > XZSlices[b.x+1][b.z].min && b.x+1 < XZSlices[b.x+1][b.z].max {
				//for _, other := range blocksOnX[b.x+1] {
				//	if b.y == other.y && b.z == other.z {
				totalX--
				//		break
			}
		}
	}
	if b.x > yzslice.min {
		if b.x-1 > XYSlices[b.x-1][b.y].min && b.x-1 < XYSlices[b.x-1][b.y].max {
			if b.x-1 > XZSlices[b.x-1][b.z].min && b.x-1 < XZSlices[b.x-1][b.z].max {
				//	for _, other := range blocksOnX[b.x-1] {
				//		if b.x > yzslice.min || (b.y == other.y && b.z == other.z) {
				totalX--
				//			break
			}
		}
	}
	if b.y < xzslice.max {
		if b.y+1 > XYSlices[b.x][b.y+1].min && b.y+1 < XYSlices[b.x][b.y+1].max {
			if b.y+1 > YZSlices[b.y+1][b.z].min && b.y+1 < YZSlices[b.y+1][b.z].max {
				//	for _, other := range blocksOnY[b.y+1] {
				//		if b.y < xzslice.max || (b.x == other.x && b.z == other.z) {
				totalY--
				//			break
			}
		}
	}
	if b.y > xzslice.min {
		if b.y-1 > XYSlices[b.x][b.y-1].min && b.y-1 < XYSlices[b.x][b.y-1].max {
			if b.y-1 > YZSlices[b.y-1][b.z].min && b.y-1 < YZSlices[b.y-1][b.z].max {
				//	for _, other := range blocksOnY[b.y-1] {
				//		if b.y > xzslice.min || (b.x == other.x && b.z == other.z) {
				totalY--
				//			break
			}
		}
	}
	if b.z < xyslice.max {
		if b.z+1 > XZSlices[b.x][b.z+1].min && b.z+1 < XZSlices[b.x][b.z+1].max {
			if b.z+1 > YZSlices[b.y][b.z+1].min && b.z+1 < YZSlices[b.y][b.z+1].max {
				//	for _, other := range blocksOnZ[b.z+1] {
				//		if b.z < xyslice.max || (b.x == other.x && b.y == other.y) {
				totalZ--
				//			break
			}
		}
	}
	if b.z > xyslice.min {
		if b.z-1 > XZSlices[b.x][b.z-1].min && b.z-1 < XZSlices[b.x][b.z-1].max {
			if b.z-1 > YZSlices[b.y][b.z-1].min && b.z-1 < YZSlices[b.y][b.z-1].max {
				//	for _, other := range blocksOnZ[b.z-1] {
				//		if b.z > xyslice.min || (b.x == other.x && b.y == other.y) {
				totalZ--
				//			break
			}
		}
	}
	return totalX + totalY + totalZ
}

var blocksOnX map[int][]block = make(map[int][]block)
var blocksOnY map[int][]block = make(map[int][]block)
var blocksOnZ map[int][]block = make(map[int][]block)
var XYSlices map[int]map[int]slice = make(map[int]map[int]slice)
var XZSlices map[int]map[int]slice = make(map[int]map[int]slice)
var YZSlices map[int]map[int]slice = make(map[int]map[int]slice)

func main() {
	allBlocks := []block{}
	file, _ := os.Open("smalltest.txt")
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
		initAndAdaptToSlice(XYSlices, x, y, z)
		initAndAdaptToSlice(XZSlices, x, z, y)
		initAndAdaptToSlice(YZSlices, y, z, x)
		allBlocks = append(allBlocks, b)
	}
	totalSides := 0
	for _, b := range allBlocks {
		totalSides += b.exposedSides()
	}
	println("total exposed sides is " + strconv.Itoa(totalSides))
}

func initAndAdaptToSlice(themap map[int]map[int]slice, a int, b int, c int) map[int]map[int]slice {
	if themap[a] == nil {
		themap[a] = make(map[int]slice)
	}
	themap[a][b] = themap[a][b].adapt(c)
	return themap
}
