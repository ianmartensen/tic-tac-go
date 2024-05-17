package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type player struct {
	symbol string
	turn   bool
}

type mark struct {
	pos   int64
	value string
	taken bool
}

type gameState struct {
	topRow    []mark
	middleRow []mark
	bottomRow []mark
}

func (s *gameState) updateGamegameState(player *player, input string, row string) {
	idx, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		log.Fatal("failed to parse int input", err)
	}
	idxPtrs := map[int64]int64{
		0: 0, 1: 1, 2: 2,
		3: 0, 4: 1, 5: 2,
		6: 0, 7: 1, 8: 2,
	}
	idx = idxPtrs[idx]

	switch row {
	case "top":
		s.topRow[idx].taken = true
		s.topRow[idx].value = player.symbol
	case "middle":
		s.middleRow[idx].taken = true
		s.middleRow[idx].value = player.symbol
	case "bottom":
		s.bottomRow[idx].taken = true
		s.bottomRow[idx].value = player.symbol
	}
}

func (s *gameState) isAvailable(row string, input string) bool {
	idx, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		log.Fatal("failed to parse int input", err)
	}

	idxPtrs := map[int64]int64{
		0: 0, 1: 1, 2: 2,
		3: 0, 4: 1, 5: 2,
		6: 0, 7: 1, 8: 2,
	}

	idx = idxPtrs[idx]

	var isTaken bool
	switch row {
	case "top":
		isTaken = s.topRow[idx].taken
	case "middle":
		isTaken = s.middleRow[idx].taken
	case "bottom":
		isTaken = s.bottomRow[idx].taken
	}

	if !isTaken {
		return true
	} else {
		return false
	}
}

func (s *gameState) isTieGame() bool {
	for i := 0; i <= 2; i++ {
		if !s.topRow[i].taken || !s.middleRow[i].taken || !s.bottomRow[i].taken {
			return false
		}
	}
	return true
}

func (s *gameState) validateWinner(input string, p player) bool {
	winConditionMap := map[string][][]int{
		// top row
		"0": {{0, 1, 2}, {0, 4, 8}, {0, 3, 6}},
		"1": {{0, 1, 2}, {1, 4, 7}},
		"2": {{0, 1, 2}, {2, 4, 6}, {2, 5, 7}},
		// middle row
		"3": {{3, 4, 5}, {0, 3, 6}},
		"4": {{0, 4, 8}, {2, 4, 6}, {1, 4, 7}, {3, 4, 5}},
		"5": {{2, 5, 8}, {3, 4, 5}},
		// bottom row
		"6": {{0, 3, 6}, {2, 4, 6}, {6, 7, 8}},
		"7": {{6, 7, 8}, {1, 4, 7}},
		"8": {{6, 7, 8}, {0, 4, 8}, {2, 5, 8}},
	}

	idxPtr := map[int]int{
		0: 0, 1: 1, 2: 2,
		3: 0, 4: 1, 5: 2,
		6: 0, 7: 1, 8: 2,
	}

	var isWinner bool
	for _, cond := range winConditionMap[input] {
		isWinner = true
		for _, pos := range cond {
			switch pos {
			case 0, 1, 2:
				position := idxPtr[pos]
				if s.topRow[position].value != p.symbol {
					isWinner = false
				}
			case 3, 4, 5:
				position := idxPtr[pos]
				if s.middleRow[position].value != p.symbol {
					isWinner = false
				}
			case 6, 7, 8:
				position := idxPtr[pos]
				if s.bottomRow[position].value != p.symbol {
					isWinner = false
				}
			}
		}

		if isWinner {
			break
		}
	}

	return isWinner
}

func renderBoard(s gameState) {
	board := ""
	board += fmt.Sprintf("_%s_|_%s_|_%s_\n", s.topRow[0].value, s.topRow[1].value, s.topRow[2].value)
	board += fmt.Sprintf("_%s_|_%s_|_%s_\n", s.middleRow[0].value, s.middleRow[1].value, s.middleRow[2].value)
	board += fmt.Sprintf(" %s | %s | %s ", s.bottomRow[0].value, s.bottomRow[1].value, s.bottomRow[2].value)

	fmt.Println(board)
}

func getCurrentPlayer(p1 *player, p2 *player) *player {
	if p1.turn && !p2.turn {
		return p1
	} else {
		return p2
	}
}

func main() {
	// init board gameState
	initGamegameState := [][]mark{
		{
			mark{pos: 0, value: "0", taken: false},
			mark{pos: 1, value: "1", taken: false},
			mark{pos: 2, value: "2", taken: false},
		},
		{
			mark{pos: 3, value: "3", taken: false},
			mark{pos: 4, value: "4", taken: false},
			mark{pos: 5, value: "5", taken: false},
		},
		{
			mark{pos: 6, value: "6", taken: false},
			mark{pos: 7, value: "7", taken: false},
			mark{pos: 8, value: "8", taken: false},
		},
	}

	gs := gameState{
		topRow:    initGamegameState[0],
		middleRow: initGamegameState[1],
		bottomRow: initGamegameState[2],
	}

	// init players
	player1 := player{symbol: "X"}
	player2 := player{symbol: "O"}

	starter := rand.Intn(2)

	if starter == 1 {
		player1.turn = true
	} else {
		player2.turn = true
	}

	// start game loop
	reader := bufio.NewReader(os.Stdin)
	for {
		renderBoard(gs)
		if player1.turn {
			fmt.Print("Player 1 turn (X): ")
		} else {
			fmt.Print("Player 2 turn (O): ")
		}
		input, err := reader.ReadString('\n')
		if err != nil {
			panic("cannot read string input")
		}

		input = strings.ToLower(strings.TrimSpace(input))

		if input == "exit" || input == "quit" || input == ":q" {
			fmt.Println("Exiting game.")
			break
		}

		currentPlayer := getCurrentPlayer(&player1, &player2)
		switchSides := false

		switch input {
		case "0", "1", "2":
			ok := gs.isAvailable("top", input)
			if ok {
				gs.updateGamegameState(currentPlayer, input, "top")
				switchSides = true
			}
		case "3", "4", "5":
			ok := gs.isAvailable("middle", input)
			if ok {
				gs.updateGamegameState(currentPlayer, input, "middle")
				switchSides = true
			}
		case "6", "7", "8":
			ok := gs.isAvailable("bottom", input)
			if ok {
				gs.updateGamegameState(currentPlayer, input, "bottom")
				switchSides = true
			}
		default:
			fmt.Println("try again -> pick a number that is not taken from 0 -> 8")
		}

		if gs.validateWinner(input, *currentPlayer) {
			renderBoard(gs)
			fmt.Println("You won!")
			break
		}

		if gs.isTieGame() {
			fmt.Println("Tie Game!")
			break
		}

		if switchSides {
			if player1.turn {
				player2.turn = true
				player1.turn = false
			} else {
				player1.turn = true
				player2.turn = false
			}
		}
	}
}
