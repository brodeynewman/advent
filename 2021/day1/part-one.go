package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	bytes := strings.Split(string(content), "\n")

	var curr int64 = 0
	var amt int = 0

	for i, num := range bytes {
		if num == "" {
			continue
		}

		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if curr == 0 {
			curr = n
			continue
		}

		if n > curr {
			// only increment if the amount increased from the previous amount
			amt++
		}

		// set the new current number
		curr = n
	}

	fmt.Println(amt)
}