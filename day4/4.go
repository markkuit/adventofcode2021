package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	numbers       []int
	boards        []board
	winningNumber int
	winningBoard  int

	lastWinning bool
)

type board struct {
	rows [5]boardRow
	won  bool
}
type boardRow [5]boardNumber
type boardNumber struct {
	number int
	marked bool
}

func readInput() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()
	str := strings.Split(s.Text(), ",")
	numbers = make([]int, len(str))
	for i, s := range str {
		numbers[i], _ = strconv.Atoi(s)
	}

	atEOF := false
out:
	for !atEOF {
		var newBoard board
		for i := 0; i < 5; {
			atEOF = !s.Scan()
			if atEOF {
				break out
			}

			line := s.Text()
			if len(line) == 0 {
				continue
			}

			var newBoardRow boardRow
			str := strings.Fields(line)
			for j, s := range str {
				newNumber, _ := strconv.Atoi(s)
				newBoardRow[j] = boardNumber{
					number: newNumber,
				}
			}

			newBoard.rows[i] = newBoardRow
			i++
		}
		boards = append(boards, newBoard)
	}
}

func drawNumber(drewNumber int) (victory bool) {
	victory = false
	for b, board := range boards {
		if board.won {
			continue
		}
		for r, row := range board.rows {
			for n, number := range row {
				if number.number == drewNumber {
					number.marked = true
					boards[b].rows[r][n] = number
				}

				if boards[b].winning() {
					winningNumber = drewNumber
					if winningBoard == 0 || lastWinning {
						winningBoard = b
					}
					victory = true
				}
			}
		}
	}
	return victory
}

func (b *board) winning() bool {
	// horizontal check
	for _, row := range b.rows {
		winningRow := true
		for _, number := range row {
			if number.marked == false {
				winningRow = false
				break
			}
		}
		if winningRow {
			b.won = true
			return true
		}
	}
	// vertical check
	for i := 0; i < 5; i++ {
		winningColumn := true
		for _, row := range b.rows {
			if row[i].marked == false {
				winningColumn = false
				break
			}
		}
		if winningColumn {
			b.won = true
			return true
		}
	}
	return false
}

func (b board) unmarkedSum() (sum int) {
	for _, row := range b.rows {
		for _, number := range row {
			if number.marked == false {
				sum += number.number
			}
		}
	}
	return sum
}

func init() {
	flag.BoolVar(&lastWinning, "lastWinning", false, "get last winning board instead")
	flag.Parse()
}

func main() {
	readInput()
	for _, n := range numbers {
		if victory := drawNumber(n); victory {
			if !lastWinning {
				break
			}
		}
	}
	fmt.Println(boards[winningBoard].unmarkedSum() * winningNumber)
}
