package main

import (
	"bufio"
	"os"
	"strconv"
)

type (
	listelement struct {
		value int
		next  *listelement
		prev  *listelement
	}
	lineelement struct {
		value   int
		listpos *listelement
	}
)

func (le *listelement) getNeighbour(x int) *listelement {
	var result *listelement = le
	for i := 0; i < x; i++ {
		result = result.next
	}
	return result
}

func (le *listelement) printline(size int) string {
	var result = ""
	var item = le
	for i := 0; i < size-1; i++ {
		result += strconv.Itoa(item.value) + ", "
		item = item.next
	}
	return result + strconv.Itoa(item.value)
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	var first *listelement
	var previous *listelement
	var theline = []lineelement{}
	var zeroelement lineelement
	var pos int
	for reader.Scan() {
		number, _ := strconv.Atoi(reader.Text())
		number *= 811589153
		element := &listelement{value: number}
		if first == nil {
			first = element
		} else {
			element.prev = previous
		}
		lel := lineelement{value: number, listpos: element}
		if number == 0 {
			zeroelement = lel
		}
		theline = append(theline, lel)
		previous = element
		pos++
	}
	first.prev = previous
	previous.next = first
	linesize := pos
	for i := linesize - 2; i >= 0; i-- {
		theline[i].listpos.next = theline[i+1].listpos
	}

	//	println(first.printline(linesize))
	//	println(zeroelement.listpos.printline(linesize))
	for ib := 0; ib < linesize*10; ib++ {
		i := ib % linesize
		if theline[i].value != 0 {
			curr := theline[i].listpos
			next := theline[i].listpos.next
			prev := theline[i].listpos.prev
			if theline[i].value > 0 {
				// move element forward
				target := theline[i].value % (linesize - 1)
				for j := 0; j < target; j++ {
					t := next.next
					prev.next = next
					next.prev = prev
					next.next = curr
					t.prev = curr
					curr.prev = next
					curr.next = t
					next = curr.next
					prev = curr.prev
				}
			} else {
				// move element backward
				target := (0 - theline[i].value) % (linesize - 1)
				for j := 0; j < target; j++ {
					t := prev.prev
					next.prev = prev
					prev.next = next
					prev.prev = curr
					t.next = curr
					curr.next = prev
					curr.prev = t
					next = curr.next
					prev = curr.prev
				}
			}
		}
		//		println(zeroelement.listpos.printline(linesize))
	}
	nr1000 := zeroelement.listpos.getNeighbour(1000 % linesize).value
	nr2000 := zeroelement.listpos.getNeighbour(2000 % linesize).value
	nr3000 := zeroelement.listpos.getNeighbour(3000 % linesize).value
	print("numbers are ")
	printints([]int{nr1000, nr2000, nr3000})
	println("sum is " + strconv.Itoa(nr1000+nr2000+nr3000))
}

func printints(line []int) {
	for i := 0; i < len(line)-1; i++ {
		print(strconv.Itoa(line[i]) + ", ")
	}
	println(line[len(line)-1])
}
