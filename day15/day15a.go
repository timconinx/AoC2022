package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

func mDistanceBetween(x1, y1, x2, y2 int) int {
	return (int)(math.Abs(float64(x1-x2)) + math.Abs((float64)(y1-y2)))
}

func main() {
	scannedIndex := 2000000
	notAtIndex := make(map[int]bool)
	beaconsAtIndex := []int{}

	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		theregexp := regexp.MustCompile(`^Sensor\sat\sx=(-?\d+),\sy=(-?\d+):\sclosest\sbeacon\sis\sat\sx=(-?\d+),\sy=(-?\d+)$`)
		thenumbers := theregexp.FindStringSubmatch(reader.Text())
		sensorx, _ := strconv.Atoi(thenumbers[1])
		sensory, _ := strconv.Atoi(thenumbers[2])
		beaconx, _ := strconv.Atoi(thenumbers[3])
		beacony, _ := strconv.Atoi(thenumbers[4])
		if sensory == scannedIndex {
			notAtIndex[sensorx] = true
		}
		if beacony == scannedIndex {
			beaconsAtIndex = append(beaconsAtIndex, beaconx)
		}
		distance := mDistanceBetween(sensorx, sensory, beaconx, beacony)
		for i := 0; i <= distance; i++ {
			othery := sensory + (distance - i)
			if othery == scannedIndex {
				for j := 0; j <= i; j++ {
					notAtIndex[sensorx-j] = true
					notAtIndex[sensorx+j] = true
				}
			}
			othery = sensory - (distance - i)
			if othery == scannedIndex {
				for j := 0; j <= i; j++ {
					notAtIndex[sensorx-j] = true
					notAtIndex[sensorx+j] = true
				}
			}
		}
	}
	for _, i := range beaconsAtIndex {
		if notAtIndex[i] {
			delete(notAtIndex, i)
		}
	}
	println("locations not containing a beacon is " + strconv.Itoa(len(notAtIndex)))
}
