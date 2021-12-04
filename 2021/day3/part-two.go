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

func multiplyPowerConsumption(g []string, e []string) int64 {
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
	return multiplyPowerConsumption(gamma, epsilon)
}

func findLargerByte(bytes []string, i int) string {
	var o int64 = 0
	var z int64 = 0

	for j := 0; j < len(bytes); j++ {
		x := string(bytes[j][i])

		inc(x, &o, &z)
	}

	if o > z || o == z {
		return "1"
	} else {
		return "0"
	}
}

func findSmallerByte(bytes []string, i int) string {
	var o int64 = 0
	var z int64 = 0

	for j := 0; j < len(bytes); j++ {
		x := string(bytes[j][i])

		inc(x, &o, &z)
	}

	if o > z || o == z {
		return "0"
	} else {
		return "1"
	}
}

func filterBytesByChar(bytes []string, bc string, index int) []string {
	blist := []string{}

	for i := 0; i < len(bytes); i++ {
		if string(bytes[i][index]) == bc {
			blist = append(blist, bytes[i])
		}
	}

	return blist
}

func calcOxygenRating(bytes []string) int64 {
	colLen := len(bytes[0])
	blist := bytes

	for i := 0; i < colLen; i++ {
		bc := findLargerByte(blist, i)
		filtered := filterBytesByChar(blist, bc, i)

		blist = filtered

		if len(blist) <= 1 {
			break;
		}
	}

	// convert the single byte list to the integer
	x, err := strconv.ParseInt(blist[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return x
}

func calcCO2Rating(bytes []string) int64 {
	colLen := len(bytes[0])
	blist := bytes

	for i := 0; i < colLen; i++ {
		bc := findSmallerByte(blist, i)
		filtered := filterBytesByChar(blist, bc, i)

		blist = filtered

		if len(blist) <= 1 {
			break;
		}
	}

	// convert the single byte list to the integer
	x, err := strconv.ParseInt(blist[0], 2, 64)
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
	bytes := strings.Split(t, "\n")

	oxygen := make(chan int64)
	co2 := make(chan int64)

	go func() {
		oxygen <- calcOxygenRating(bytes)
	}()

	go func() {
		co2 <- calcCO2Rating(bytes)
	}()

	or := <- oxygen
	cr := <- co2

	// multiply oxygen rating by co2 rating
	fmt.Println(or * cr)
}