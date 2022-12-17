package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	valve struct {
		id      string
		rate    int
		leadsto []string
	}
	valvepair struct {
		from string
		to   string
	}
)

var valves = make(map[string]valve)
var openablevalves int
var vdistances map[string]map[string]int = make(map[string]map[string]int)

func main() {
	var distances map[valvepair]int = make(map[valvepair]int)
	var dexist map[valvepair]bool = make(map[valvepair]bool)
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		parts := regexp.MustCompile(`^Valve\s(\w+)\shas\sflow\srate=(\d+);`).FindStringSubmatch(text)
		valveid := parts[1]
		flowrate, _ := strconv.Atoi(parts[2])
		if flowrate > 0 {
			openablevalves++
		}
		others := regexp.MustCompile(`valve(s?)\s(.+)$`).FindStringSubmatch(text)[2]
		thisvalve := valve{id: valveid, rate: flowrate}
		thisvalve.leadsto = strings.Split(others, ", ")
		for _, to := range thisvalve.leadsto {
			vp1 := valvepair{from: valveid, to: to}
			vp2 := valvepair{from: to, to: valveid}
			distances[vp1] = 1
			dexist[vp1] = true
			distances[vp2] = 1
			dexist[vp2] = true
		}
		valves[valveid] = thisvalve
	}
	for {
		if len(distances) < len(valves)*len(valves) {
			for pair, distance := range distances {
				for _, to := range valves[pair.to].leadsto {
					vp1 := valvepair{from: pair.from, to: to}
					vp2 := valvepair{from: to, to: pair.from}
					if !dexist[vp1] {
						distances[vp1] = distance + 1
						dexist[vp1] = true
					} else if distance+1 < distances[vp1] {
						distances[vp1] = distance + 1
					}
					if !dexist[vp2] {
						distances[vp2] = distance + 1
						dexist[vp2] = true
					} else if distance+1 < distances[vp2] {
						distances[vp2] = distance + 1
					}
				}
			}
		} else {
			break
		}
	}
	for _, v := range valves {
		vd := make(map[string]int)
		for _, v2 := range valves {
			if v.id != v2.id && v2.rate > 0 {
				vd[v2.id] = distances[valvepair{from: v.id, to: v2.id}]
			}
		}
		vdistances[v.id] = vd
	}
	println("" + strconv.Itoa(len(distances)) + " number of distances")
	/*println("distances from AA")
	for target, distance := range vdistances["AA"] {
		println(target + " at distance " + strconv.Itoa(distance))
	}*/
	maxpressure := maximumPressureFrom("AA", 30, 0, []string{})
	println("maximum pressure deflated is " + strconv.Itoa(maxpressure))
}

func maximumPressureFrom(vid string, time int, openvalves int, opened []string) int {
	if openvalves == openablevalves {
		return 0
	}
	var maxp int
	for target, distance := range vdistances[vid] {
		timeworth := time - (distance + 1)
		if timeworth > 0 && !contains(opened, target) {
			p := timeworth*valves[target].rate + maximumPressureFrom(target, timeworth, openvalves+1, append(opened, target))
			if p > maxp {
				maxp = p
			}
		}
	}
	return maxp
}

func printPath(path []string) string {
	result := "["
	for i := 0; i < len(path); i++ {
		result += path[i]
		if i != len(path)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}

func contains(group []string, a string) bool {
	for _, g := range group {
		if g == a {
			return true
		}
	}
	return false
}
