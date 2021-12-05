package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type Piece struct {
	marked bool
	value  string
}

type Board struct {
	layout [  ][]Piece
	won      bool
	id       int
	winPiece string
}

func newBoard(id int) Board {
	b := Board{layout: [][]Piece{}, id: id}

	return b
}

func newPiece(v string) Piece {
	p := Piece{value: string(v), marked: false}
	
	return p
}

func fillBoard(b *Board, layout []Piece) *Board {
	b.layout = append(b.layout, layout)

	return b
}

func buildGame(boards *[]Board, bytes []string) {
	rc := 0
	b := newBoard(len(*boards) + 1)

	// start at 1 since we already have header captured
	for i := 1; i < len(bytes); i++ {
		trimmed := strings.TrimSpace(string(bytes[i]))

		if len(trimmed) < 1 {
			continue;
		}

		split := strings.Split(trimmed, " ")

		if (rc == 5) {
			*boards = append(*boards, b)
			rc = 0
			id := len(*boards) + 1

			// reset our board since we're building a new one
			b = newBoard(id)
		}

		layout := []Piece{}

		for j := 0; j < len(split); j++ {
			// check if botched spaces
			if len(split[j]) > 0 {
				p := newPiece(split[j])

				layout = append(layout, p)
			}
		}

		rc++
		fillBoard(&b, layout)

		// make sure to fill our last board at the end of our chunking
		if i == (len(bytes) - 1) {
			*boards = append(*boards, b)
		}
	}
}

func (p *Piece) checkPiece(val string) {
	if p.value == val {
		(*p).marked = true
	}
}

func (b *Board) addMatch(val string) {
	for i := 0; i < len(b.layout); i++ {
		row := b.layout[i]

		for j := 0; j < len(row); j++ {
			p := &row[j]
			p.checkPiece(val)
		}
	}
}

func (b *Board) checkWinner() {
	// we know the bingo boards will always be 5x5
	for i := 0; i < 5; i++ {
		rWin := 0;

		for j := 0; j < 5; j++ {
			p := b.layout[i][j]

			if (p.marked == true) {
				rWin++
			}
		}

		if (rWin == 5) {
			(*b).won = true
			break;
		}
	}

	// no need to coninue looking at columns if a row won
	if (*b).won {
		return
	}

	for i := 0; i < 5; i++ {
		cWin := 0;

		// check column
		for j := 0; j < 5; j++ {
			col := b.layout[j]
			p := col[i]

			if (p.marked == true) {
				cWin++
			}
		}

		if (cWin == 5) {
			(*b).won = true
			break;
		}
	}
}

func calcWinningAmount(b *Board) int64 {
	var x int64 = 0
	y, err := strconv.ParseInt(b.winPiece, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(b.layout); i++ {
		for j := 0; j < len(b.layout[i]); j++ {
			p := b.layout[i][j]

			if (p.marked == false) {
				num, err := strconv.ParseInt(p.value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				x += num
			}
		}
	}
	
	return x * y
}

func (b *Board) storeWinPiece(val string) {
	(*b).winPiece = val
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	t := strings.TrimSpace(string(content))
	bytes := strings.Split(t, "\n")
	readings := strings.Split(bytes[0], ",")

	var boards []Board
	var winners []*Board
	wonBoards := make(map[int]bool)

	buildGame(&boards, bytes)

	for i := 0; i < len(readings) && len(boards) > 0; i++ {
		val := string(readings[i])

		// scan our boards and fill in matches
		for j := 0; j < len(boards); j++ {
			// allows us to skip checking a board that already won
			if (wonBoards[boards[j].id] == true) {
				continue;
			}

			b := boards[j]
			b.addMatch(val)
		}

		// once 5 checks have been done, see if there are any winners
		if (i < 5) {
			continue;
		}

		for j := 0; j < len(boards); j++ {
			// allows us to skip checking a board that already won
			if (wonBoards[boards[j].id] == true) {
				continue;
			}

			b := &boards[j]
			b.checkWinner()

			if (b.won) {
				b.storeWinPiece(val)
				wonBoards[boards[j].id] = true

				// append our new winner to the end of the winners list
				winners = append(winners, b)
			}
		}
	}

	// calc our last winner's amount
	lastWinAmt := calcWinningAmount(winners[len(winners) - 1])
	fmt.Println(lastWinAmt)
}