package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var theline = []string{}
var positions = make(map[string]int)
var exists = make(map[string]bool)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	pos := 0
	for reader.Scan() {
		number := uniqueString(reader.Text())
		theline = append(theline, number)
		positions[number] = pos
		exists[number] = true
		pos++
	}
	//	printconga(theline)
	linesize := pos
	conga := make([]string, len(theline))
	copy(conga, theline)
	for _, ns := range theline {
		n := fromUniqueString(ns)
		oldpos := positions[ns]
		var newpos int
		if n > 0 {
			newpos = (positions[ns] + (n % (linesize - 1))) % linesize
		} else {
			newpos = (positions[ns] + n) % linesize
		}
		if newpos <= 0 {
			newpos = linesize + newpos - 1
		}
		if positions[ns]+(n%(linesize-1)) > linesize {
			newpos = (newpos + 1) % linesize
		}
		if n != 0 {
			if oldpos < newpos {
				for {
					opp := (oldpos + 1) % linesize
					conga[oldpos] = conga[opp]
					positions[conga[oldpos]] = oldpos
					oldpos = opp
					if oldpos == newpos {
						break
					}
				}
			} else if oldpos > newpos {
				for {
					conga[oldpos] = conga[oldpos-1]
					positions[conga[oldpos]] = oldpos
					oldpos--
					if oldpos == newpos {
						break
					}
				}
			}
			conga[newpos] = ns
			positions[ns] = newpos
		}
		//		printconga(conga)
	}
	nr1000 := fromUniqueString(conga[(positions["0"]+1000)%linesize])
	nr2000 := fromUniqueString(conga[(positions["0"]+2000)%linesize])
	nr3000 := fromUniqueString(conga[(positions["0"]+3000)%linesize])
	print("numbers are ")
	printconga([]string{strconv.Itoa(nr1000), strconv.Itoa(nr2000), strconv.Itoa(nr3000)})
	println("sum is " + strconv.Itoa(nr1000+nr2000+nr3000))
}

func printconga(line []string) {
	for i := 0; i < len(line)-1; i++ {
		print(line[i] + ", ")
	}
	println(line[len(line)-1])
}

func uniqueString(s string) string {
	var suffix string
	if exists[s] {
		var i int
		for {
			suffix = ":" + strconv.Itoa(i)
			if !exists[s+suffix] {
				return s + suffix
			}
			i++
		}
	}
	return s
}

func fromUniqueString(s string) int {
	result, err := strconv.Atoi(strings.Split(s, ":")[0])
	if err != nil {
		panic(err)
	}
	return result
}
