package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func mDistanceBetween(x1, y1, x2, y2 int) int {
	return (int)(math.Abs(float64(x1-x2)) + math.Abs((float64)(y1-y2)))
}

func main() {
	startTime := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	theInput := [][]int{}
	for reader.Scan() {
		theregexp := regexp.MustCompile(`^Sensor\sat\sx=(-?\d+),\sy=(-?\d+):\sclosest\sbeacon\sis\sat\sx=(-?\d+),\sy=(-?\d+)$`)
		thenumbers := theregexp.FindStringSubmatch(reader.Text())
		sensorx, _ := strconv.Atoi(thenumbers[1])
		sensory, _ := strconv.Atoi(thenumbers[2])
		beaconx, _ := strconv.Atoi(thenumbers[3])
		beacony, _ := strconv.Atoi(thenumbers[4])
		distance := mDistanceBetween(sensorx, sensory, beaconx, beacony)

		theInput = append(theInput, []int{sensorx, sensory, distance})
	}

	//bound := 20
	bound := 4000000
	for scannedIndex := 0; scannedIndex < bound; scannedIndex++ {
		notAtIndex := make(map[int]bool)
		notHere := false
		for _, input := range theInput {
			sensorx := input[0]
			sensory := input[1]
			distance := input[2]

			if sensory == scannedIndex {
				notAtIndex[sensorx] = true
			}
			for i := 0; i <= distance; i++ {
				othery := sensory + (distance - i)
				if othery == scannedIndex {
					for j := 0; j <= i; j++ {
						if sensorx-j >= 0 {
							notAtIndex[sensorx-j] = true
						}
						if sensorx+j <= bound {
							notAtIndex[sensorx+j] = true
						}
					}
				}
				othery = sensory - (distance - i)
				if othery == scannedIndex {
					for j := 0; j <= i; j++ {
						if sensorx-j >= 0 {
							notAtIndex[sensorx-j] = true
						}
						if sensorx+j <= bound {
							notAtIndex[sensorx+j] = true
						}
					}
				}
				if len(notAtIndex) > bound {
					notHere = true
					break
				}
			}
			if notHere {
				break
			}
		}
		if !notHere {
			println("found at y " + strconv.Itoa(scannedIndex))
			for i := 0; i < bound; i++ {
				if !notAtIndex[i] {
					println("found at x " + strconv.Itoa(i))
					tuningfrequency := i*4000000 + scannedIndex
					println("tuning frequency is " + strconv.Itoa(tuningfrequency))
				}
			}
		}
	}
	elapsed := time.Since(startTime)
	println("elapsed time: " + elapsed.String())
}
