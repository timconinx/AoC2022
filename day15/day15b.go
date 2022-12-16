package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type linerange struct {
	from int
	to   int
}

func (lr linerange) Equals(other linerange) bool {
	return lr.from == other.from && lr.to == other.to
}

func (lr linerange) String() string {
	return "[" + strconv.Itoa(lr.from) + "->" + strconv.Itoa(lr.to) + "]"
}

func (lr linerange) add(lr2 linerange) (linerange, linerange, bool) {
	if lr.from <= lr2.from && lr.to >= lr2.to {
		return lr, lr, true
	}
	if lr2.from <= lr.from && lr2.to >= lr.to {
		return lr2, lr2, true
	}
	if lr.to >= lr2.from && lr.from <= lr2.to {
		var from, to int
		if lr.from <= lr2.from {
			from = lr.from
		} else {
			from = lr2.from
		}
		if lr.to >= lr2.to {
			to = lr.to
		} else {
			to = lr2.to
		}
		return linerange{from: from, to: to}, linerange{}, true
	}
	if lr.from <= lr2.to && lr.to >= lr2.from {
		var from, to int
		if lr.from <= lr2.from {
			from = lr.from
		} else {
			from = lr2.from
		}
		if lr.to >= lr2.to {
			to = lr.to
		} else {
			to = lr2.to
		}
		return linerange{from: from, to: to}, linerange{}, true
	}
	if lr.from == lr2.to+1 || lr2.from == lr.to+1 {
		var from, to int
		if lr.from <= lr2.from {
			from = lr.from
		} else {
			from = lr2.from
		}
		if lr.to >= lr2.to {
			to = lr.to
		} else {
			to = lr2.to
		}
		return linerange{from: from, to: to}, linerange{}, true
	}
	return lr, lr2, false
}

func (lr linerange) mergeWith(ranges []linerange) []linerange {
	result := []linerange{}
	if len(ranges) > 0 {
		for {
			var i int
			if l1, l2, success := lr.add(ranges[i]); success {
				result = append(result, l1)
				break
			} else {
				result = append(result, []linerange{l1, l2}...)
			}
		}
	}
	return result
}

func mDistanceBetween(x1, y1, x2, y2 int) int {
	return (int)(math.Abs(float64(x1-x2)) + math.Abs((float64)(y1-y2)))
}

func main() {
	startTime := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	//stdinputreader := bufio.NewScanner(os.Stdin)
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
	fullrange := linerange{from: 0, to: bound}
	for scannedIndex := 0; scannedIndex < bound; scannedIndex++ {
		ranges := []linerange{}
		for _, input := range theInput {
			sensorx := input[0]
			sensory := input[1]
			distance := input[2]

			if sensory == scannedIndex {
				if sensorx >= 0 && sensorx <= bound {
					ranges = append(ranges, linerange{from: sensorx, to: sensorx})
				}
			}
			for i := 0; i <= distance; i++ {
				othery := sensory + (distance - i)
				if othery == scannedIndex {
					lr := linerange{from: sensorx - i, to: sensorx + i}
					if lr.from < 0 {
						lr.from = 0
					}
					if lr.to > bound {
						lr.to = bound
					}
					ranges = append(ranges, lr)
				}
				othery = sensory - (distance - i)
				if othery == scannedIndex {
					lr := linerange{from: sensorx - i, to: sensorx + i}
					if lr.from < 0 {
						lr.from = 0
					}
					if lr.to > bound {
						lr.to = bound
					}
					ranges = append(ranges, lr)
				}
			}
		}
		/*if scannedIndex%100 == 0 {
			print(".")
		}
		if scannedIndex%4000 == 0 {
			println(".")
		}*/
		//println("cleaning ranges for line " + strconv.Itoa(scannedIndex) + ": nrranges is " + strconv.Itoa(len(ranges)))
		for {
			/*if scannedIndex == 12 {
				for _, lr := range ranges {
					println(lr.String())
				}
				println("enter for iteration")
				stdinputreader.Scan()
			}*/
			if len(ranges) <= 2 {
				break
			}
			newranges := []linerange{}
			baserange := ranges[0]
			for i := 1; i < len(ranges); i++ {
				if l1, l2, success := baserange.add(ranges[i]); success {
					baserange = l1
					if baserange.Equals(fullrange) {
						ranges = []linerange{fullrange}
						break
					}
				} else {
					newranges = append(newranges, l2)
				}
			}
			newranges = append(newranges, baserange)
			//newranges = append([]linerange{baserange}, newranges...)
			ranges = newranges
		}
		if len(ranges) > 1 {
			if _, _, success := ranges[0].add(ranges[1]); !success {
				println("found at y " + strconv.Itoa(scannedIndex))
				var x int
				if ranges[0].from == 0 {
					x = ranges[0].to + 1
				} else {
					x = ranges[0].from - 1
				}
				println("found at x " + strconv.Itoa(x))
				tuningfrequency := x*4000000 + scannedIndex
				println("tuning frequency is " + strconv.Itoa(tuningfrequency))
				break
			}
		}
	}
	elapsed := time.Since(startTime)
	println("elapsed time: " + elapsed.String())
}
