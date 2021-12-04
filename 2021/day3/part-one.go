package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

const One = "1"
const Zero = "0"

func inc(x string, one *int64, zero *int64) {
	if x == One {
		*one++
	}

	if x == Zero {
		*zero++
	}
}

func multiply(g []string, e []string) int64 {
	x := strings.Join(g[:], "")
	y := strings.Join(e[:], "")

	xi, xErr := strconv.ParseInt(x, 2, 64)  
	yi, yErr := strconv.ParseInt(y, 2, 64)

	if xErr != nil { log.Fatal(xErr) }
	if yErr != nil { log.Fatal(yErr) }

	return xi * yi
}

func calcPowerConsumption(bytes []string) int64 {
	colLen := len(bytes[0])

	gamma := []string{}
	epsilon := []string{}

	for i := 0; i < colLen; i++ {
		var o int64 = 0
		var z int64 = 0

		for j := 0; j < len(bytes); j++ {
			x := string(bytes[j][i])

			inc(x, &o, &z)
		}

		if o > z {
			epsilon = append(epsilon, "0")
			gamma = append(gamma, "1")
		} else {
			epsilon = append(epsilon, "1")
			gamma = append(gamma, "0")
		}
	}

	// multiply gamma & episol binary number value
	return multiply(gamma, epsilon)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	t := strings.TrimSpace(string(content))
	bytes := strings.Split(t, "\n")

	var pc int64 = calcPowerConsumption(bytes)

	fmt.Println(pc)
}