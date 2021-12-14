package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	last      int
	inputs    []int
	increases = -1
)

func readInput() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		value, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, value)
	}
}

func main() {
	readInput()
	for i := 0; i+2 < len(inputs); i++ {
		value := inputs[i] + inputs[i+1] + inputs[i+2]
		if value > last {
			increases++
		}
		last = value
	}
	fmt.Println(increases)
}
