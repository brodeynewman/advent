package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type Point struct {
	pos []int64
	hits int
}

func newPoint(x int64, y int64) Point {
	p := Point{hits: 0, pos: []int64{x, y}}

	return p
}

func findMaxAxis(p [][][]string) int64 {
	var sb strings.Builder

	for i := 0; i < len(p); i++ {
		p := p[i]

		p1 := p[0]
		p2 := p[1]

		axisLen := sb.Len()

		if len(p1[0]) > axisLen || len(p1[0]) > axisLen || len(p2[0]) > axisLen || len(p2[1]) > axisLen {
			sb.WriteString("9")
		}
	}

	x, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return x
}

func buildPlot(x int64) [][]Point {
	var plot [][]Point

	for i := int64(0); i <= x; i++ {
		var row []Point

		for j := int64(0); j <= x; j++ {
			p := newPoint(i, j)

			row = append(row, p)
		}

		plot = append(plot, row)
	}

	return plot
}

func buildPointPairs() [][][]string {
	var points [][][]string

	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	t := strings.TrimSpace(string(content))
	split := strings.Split(t, "\n")

	for i := 0; i < len(split); i++ {
		s := strings.Split(split[i], "->")
		p1 := strings.Split(strings.TrimSpace(s[0]), ",")
		p2 := strings.Split(strings.TrimSpace(s[1]), ",")

		x1 := p1[0]
		x2 := p2[0]

		y1 := p1[1]
		y2 := p2[1]

		px := []string{x1, y1}
		py := []string{x2, y2}
		ppair := [][]string{px, py}

		points = append(points, ppair)
	}

	return points
}

func parseInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return x
}

func markPlot(plot *[][]Point, pairs [][][]string) {
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]

		p1 := pair[0]
		p2 := pair[1]

		x1 := parseInt64(p1[0])
		x2 := parseInt64(p2[0])

		y1 := parseInt64(p1[1])
		y2 := parseInt64(p2[1])

		if x1 > x2 {
			xdiff := x1 - x2

			fmt.Println(xdiff)
		}

		if x2 > x1 {
			xdiff := x2 - x1

			fmt.Println(xdiff)
		}

		if y2 > y1 {
			ydiff := y2 - y1

			fmt.Println(ydiff)
		}

		if y1 > y2 {
			ydiff := y1 - y2

			fmt.Println(ydiff)
		}

		fmt.Println(x1, x2, y1, y2)
	}
}

func main() {
	pairs := buildPointPairs()
	maxLen := findMaxAxis(pairs)
	plot := buildPlot(maxLen)

	markPlot(&plot, pairs)

	fmt.Println(maxLen, plot)
}