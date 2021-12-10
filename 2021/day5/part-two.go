package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func parseInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return x
}

func getDiagonalPoints(x1 string, y1 string, x2 string, y2 string) []string {
	var plist []string

	var xi int64 = parseInt64(x1)
	var yi int64 = parseInt64(y1)

	var xMove int64
	var yMove int64

	if parseInt64(x1) < parseInt64(x2) {
		xMove = 1
	} else {
		xMove = -1
	}

	if parseInt64(y1) < parseInt64(y2) {
		yMove = 1
	} else {
		yMove = -1
	}

	var x int64 = xi;
	var y int64 = yi;

	for true {
		var point string = strconv.FormatInt(x, 10) + "," + strconv.FormatInt(y, 10)
		plist = append(plist, point)

		if x == parseInt64(x2) {
			break
		}

		x += xMove
		y += yMove
	}

	return plist
}

func buildPoints(p [][][]string) []string {
	var points []string

	for i := 0; i < len(p); i++ {
		pair := p[i]
		x1 := string(pair[0][0])
		y1 := string(pair[0][1])

		x2 := string(pair[1][0])
		y2 := string(pair[1][1])

		if (x1 == x2) {	
			py1 := parseInt64(y1)
			py2 := parseInt64(y2)

			var plist []string
			
			if py1 < py2 {
				for i := py1; i <= py2; i++ {
					y := strconv.FormatInt(i, 10)
					var point string = x1 + "," + y

					plist = append(plist, point)
				}
			} else {
				for i := py2; i <= py1; i++ {
					y := strconv.FormatInt(i, 10)
					var point string = x1 + "," + y

					plist = append(plist, point)
				}
			}

			points = append(points, plist...)
		}

		if (y1 == y2) {
			px1 := parseInt64(x1)
			px2 := parseInt64(x2)

			var plist []string
			
			if px1 < px2 {
				for i := px1; i <= px2; i++ {
					x := strconv.FormatInt(i, 10)
					var point string = x + "," + y2
	
					plist = append(plist, point)
				}
			} else {
				for i := px2; i <= px1; i++ {
					x := strconv.FormatInt(i, 10)
					var point string = x + "," + y2
	
					plist = append(plist, point)
				}
			}

			points = append(points, plist...)
		}

		// diagonal lines
		if (x1 != x2 && y1 != y2) {
			plist := getDiagonalPoints(x1, y1, x2, y2)
			points = append(points, plist...)
		}
	}

	return points
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

		var s1 string = x1
		var s2 string = x2
		var s3 string = y1
		var s4 string = y2

		a := []string{s1, s3}
		b := []string{s2, s4}

		ppair := [][]string{a, b}
		points = append(points, ppair)
	}

	return points
}

func main() {
	cache := make(map[string]int64)
	pairs := buildPointPairs()

	points := buildPoints(pairs)

	for i := 0; i < len(points); i++ {
		if cache[points[i]] >= 1 {
			cache[points[i]]++
		} else {
			cache[points[i]] = 1
		}
	}

	var dp int64 = 0

	for _, hits := range cache {
		if hits > 1 {
			dp++
		}
	}

	fmt.Println(dp)
}