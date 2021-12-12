package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

const NumDays = 256

func parseInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return x
}

type Key struct {
	daysLeft int
	count int
}

func calcFuel(x int, y int) int {
	m := 0
	s := 0

	var inc int
	var final int

	if x < y {
		inc = x
		final = y
	} else {
		inc = y
		final = x
	}

	for i := inc; i < final; i++ {
		m = m + s + 1
		s++
	}

	return m
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	t := strings.TrimSpace(string(content))
	inputs := strings.Split(t, ",")

	var positions []int64

	var cache map[int]int = make(map[int]int)

	for i := 0; i < len(inputs); i++ {
		positions = append(positions, parseInt64(inputs[i]))
	}

	var lowest int = -1
	var tp int

	for i := 1; i <= len(positions); i++ {
		pos := i
		n := 0

		for j := 0; j < len(positions); j++ {
			fuel := calcFuel(pos, int(positions[j]))

			n += fuel
		}

		cache[int(pos)] = n
	}

	for pos, fuel := range cache {
		if lowest == -1 || fuel < lowest {
			tp = pos
			lowest = fuel
		}
	}

	fmt.Println(tp, lowest)
}