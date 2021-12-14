package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pos, depth, aim int

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		command, unitsString := fields[0], fields[1]
		units, _ := strconv.Atoi(unitsString)

		switch command {
		case "forward":
			pos += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	fmt.Println(pos * depth)
}
