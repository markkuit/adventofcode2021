package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	gammaString, epsilonString string
	inputs                     []string
)

func readInput() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		inputs = append(inputs, s.Text())
	}
}

func main() {
	readInput()

	for i := 0; i < len(inputs[0]); i++ { // we're correct in assuming every input has the same len, so we just take any
		var zeroes, ones int
		for _, value := range inputs {
			switch string(value[i]) {
			case "0":
				zeroes++
			case "1":
				ones++
			}
		}
		if zeroes > ones {
			gammaString += "0"
			epsilonString += "1"
		} else {
			gammaString += "1"
			epsilonString += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gamma * epsilon)
}
