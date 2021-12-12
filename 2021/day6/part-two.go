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

func deduceSpawnedFishRescursively(count int, daysLeft int, cache *map[Key]int) int {
	ndl := daysLeft - (count + 1)

	if ndl < 0 { return 1 }

	if _, ok := (*cache)[Key{ndl, count}]; !ok {
		v := deduceSpawnedFishRescursively(8, ndl, cache) + deduceSpawnedFishRescursively(6, ndl, cache)
		(*cache)[Key{ndl, count}] = v
	}

	x := (*cache)[Key{ndl, count}]

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

	var numFish int = 0
	var cache map[Key]int = make(map[Key]int)

	for j := 0; j < len(fish); j++ {
		count := fish[j]

		numFish += deduceSpawnedFishRescursively(int(count), NumDays, &cache)
	}

	fmt.Println(numFish)
}