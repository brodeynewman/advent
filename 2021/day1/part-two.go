package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func addSlidingWindow(n1s string, n2s string, n3s string) int64 {
	n1, err1 := strconv.ParseInt(n1s, 10, 64)
	n2, err2 := strconv.ParseInt(n2s, 10, 64)
	n3, err3 := strconv.ParseInt(n3s, 10, 64)

	if err1 != nil { log.Fatal(err1) }
	if err2 != nil { log.Fatal(err2) }
	if err3 != nil { log.Fatal(err3) }

	return n1 + n2 + n3
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	bytes := strings.Split(string(content), "\n")
	length := len(bytes)

	var curr int64 = 0
	var amt int = 0

	for i := 0; i < length; i++ {
		if (i + 2 >= length) {
			break
		}

		n1 := bytes[i]
		n2 := bytes[i + 1]
		n3 := bytes[i + 2]

		if n1 == "" || n2 == "" || n3 == "" {
			continue
		}

		window := addSlidingWindow(n1, n2, n3)

		if curr == 0 {
			curr = window
			continue
		}

		if window > curr {
			// only increment if the amount increased from the previous amount
			amt++
		}

		// set the new current number
		curr = window
	}

	fmt.Println(amt)
}