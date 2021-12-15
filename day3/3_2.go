package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	inputs []string
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

func determineRating(numbers []string, reverse bool) int64 {
	for i := 0; len(numbers) > 1 && i < len(numbers[0]); i++ {
		var kept []string
		var zeroes, ones int
		var keep byte

		for _, value := range numbers {
			switch value[i] {
			case '0':
				zeroes++
			case '1':
				ones++
			}
		}

		if reverse {
			zeroes, ones = ones, zeroes
		}
		if zeroes > ones {
			keep = '0'
		} else if ones > zeroes {
			keep = '1'
		} else {
			if !reverse {
				keep = '1'
			} else {
				keep = '0'
			}
		}

		for _, value := range numbers {
			if value[i] == keep {
				kept = append(kept, value)
			}
		}
		numbers = make([]string, len(kept))
		copy(numbers, kept)
	}

	dec, err := strconv.ParseInt(numbers[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return dec
}

func main() {
	readInput()
	oxygenRating := determineRating(inputs, false)
	co2Rating := determineRating(inputs, true)
	fmt.Println(oxygenRating * co2Rating)
}
