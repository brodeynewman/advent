package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

const Up string = "up"
const Down string = "down"
const Forward string = "forward"

func unpack(cmd string) (string, int64) {
	l := strings.Split(string(cmd), " ")
	n, err := strconv.ParseInt(l[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return l[0], n
}

func doCommand(cmd string, amt int64, a *int64, h *int64, d *int64) {
	if cmd == Forward {
		*h += amt
		*d += *a * amt
	}

	if cmd == Up {
		*a -= amt
	}

	if cmd == Down {
		*a += amt
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	bytes := strings.Split(string(content), "\n")

	var h int64 = 0
	var d int64 = 0
	var a int64 = 0

	for _, cmd := range bytes {
		if cmd == "" {
			continue
		}

		command, x := unpack(cmd)
		doCommand(command, x, &a, &h, &d)
	}

	// multiply height * depth to get our answer.
	// aim is already included here.
	fmt.Println(h * d)
}