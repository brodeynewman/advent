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

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	t := strings.TrimSpace(string(content))
	inputs := strings.Split(t, ",")

	var fish []int64
	
	for i := 0; i < len(inputs); i++ {
		fish = append(fish, parseInt64(inputs[i]))
	}

	for i := 0; i < NumDays; i++ {
		fmt.Println(i)

		for i := 0; i < len(fish); i++ {
			if (fish[i] == 0) {
				fish[i] = 7
				fish = append(fish, 9)
			}

			x := fish[i] - 1
			fish[i] = x
		}
	}

	fmt.Println(len(fish))
}