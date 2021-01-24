package main

import (
	"fmt"
	"math"
	"strings"
)

type Move struct {
	x, y   int
	player string
}

type MoveValue struct {
	Move
	value int
}

type TicTacToe [][]string

func (t *TicTacToe) Init(n int) {
	*t = make([][]string, n)
	for i := range *t {
		(*t)[i] = make([]string, n)
		for j := range (*t)[i] {
			(*t)[i][j] = "_"
		}
	}
}

func (t *TicTacToe) Print() {
	for i := 0; i < len(*t); i++ {
		fmt.Printf("%s\n", strings.Join((*t)[i], " "))
	}
}

func (t *TicTacToe) Play(move Move) error {
	n := len(*t)
	if move.x <= 0 || move.x > n || move.y <= 0 || move.y > n {
		return fmt.Errorf("error: Invalid index")
	}
	if (*t)[move.x-1][move.y-1] != "_" {
		return fmt.Errorf("error: Square not empty")
	}
	(*t)[move.x-1][move.y-1] = move.player
	return nil
}

func (t *TicTacToe) isRow(n int) (bool, string) {
	row := (*t)[n]
	player := row[0]
	if player == "_" {
		return false, player
	}
	for i := range row {
		if row[i] != player {
			return false, player
		}
	}
	return true, player
}

func (t *TicTacToe) isCol(n int) (bool, string) {
	player := (*t)[0][n]
	if player == "_" {
		return false, player
	}
	for _, row := range *t {
		if row[n] != player {
			return false, player
		}
	}
	return true, player
}

func (t *TicTacToe) isDiagonal(left bool) (bool, string) {
	x, dX := 1, 1
	if !left {
		x, dX = len(*t), -1
	}
	player := (*t)[0][x-1]
	if player == "_" {
		return false, player
	}
	for i, row := range *t {
		if row[x+dX*i-1] != player {
			return false, player
		}
	}
	return true, player
}

func (t *TicTacToe) IsWinning() (bool, string) {
	n := len(*t)
	for i := 0; i < n; i++ {
		if finished, winner := t.isRow(i); finished {
			return finished, winner
		}
		if finished, winner := t.isCol(i); finished {
			return finished, winner
		}
	}
	if finished, winner := t.isDiagonal(true); finished {
		return finished, winner
	}
	if finished, winner := t.isDiagonal(false); finished {
		return finished, winner
	}
	return false, ""
}

func (t *TicTacToe) AreEmptyCells() bool {
	for i, row := range *t {
		for j := range row {
			if (*t)[i][j] == "_" {
				return true
			}
		}
	}
	return false
}

func (t *TicTacToe) ValueForPlayer(player string) int {
	isWinning, winner := t.IsWinning()
	if isWinning {
		if winner == player {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func (t *TicTacToe) GetMoves(player string) (moves []Move) {
	for i, row := range *t {
		for j := range row {
			if (*t)[i][j] == "_" {
				moves = append(moves, Move{i + 1, j + 1, player})
			}
		}
	}
	return
}

func (t *TicTacToe) Copy() TicTacToe {
	newCopy := TicTacToe{}
	newCopy.Init(len(*t))
	for i, row := range *t {
		for j := range row {
			newCopy[i][j] = (*t)[i][j]
		}
	}
	return newCopy
}

func max(x, y MoveValue) MoveValue {
	if x.value >= y.value {
		return x
	}
	return y
}
func min(x, y MoveValue) MoveValue {
	if x.value < y.value {
		return x
	}
	return y
}

func (t *TicTacToe) GetPositionValue(maxDepth int, player string) (bestValue int) {
	maximizingPlayer := player == "O"
	moves := t.GetMoves(player)
	isTerminal, _ := t.IsWinning()
	if maxDepth == 0 || isTerminal || len(moves) == 0 {
		//fmt.Printf("Terminal: %s -> %d", move, position.ValueForPlayer(move.player))
		return t.ValueForPlayer("O") // return Move zero value
	}
	if maximizingPlayer {
		bestValue = math.MinInt32
		for _, mv := range moves {
			newPos := t.Copy()
			err := newPos.Play(mv)
			if err != nil {
				panic(err)
			}
			val := newPos.GetPositionValue(maxDepth-1, "X")
			if val > bestValue {
				bestValue = val
			}
		}
	} else { // minimizing player
		bestValue = math.MaxInt32
		for _, mv := range moves {
			newPos := t.Copy()
			err := newPos.Play(mv)
			if err != nil {
				panic(err)
			}
			val := newPos.GetPositionValue(maxDepth-1, "O")
			if val < bestValue {
				bestValue = val
			}
		}
	}
	t.Print()
	fmt.Printf("Player: %s, Value: %d\n", player, bestValue)
	return
}

var ttt TicTacToe

func main() {
	n := 3
	ttt.Init(n)
	ttt.Print()
	player := "X"
	var winner string
	var move Move
	for finished := false; !finished && ttt.AreEmptyCells(); finished, winner = ttt.IsWinning() {
		if player == "O" { // computer's turn
			moves := ttt.GetMoves(player)
			maxValue := math.MinInt32
			for _, mv := range moves {
				newPos := ttt.Copy()
				err := newPos.Play(mv)
				if err != nil {
					panic(err)
				}
				value := newPos.GetPositionValue(10, "O")
				fmt.Printf("Move: %v -> %d\n", mv, value)
				if value >= maxValue {
					maxValue = value
					move = mv
				}
			}
			fmt.Printf("Computer plays: '%v', value: %d\n", move, maxValue)
		} else { // your turn - enter move
			fmt.Printf("Enter X Y for '%s'\n", player)
			var x, y int
			//var str string
			l, err := fmt.Scanf("%d %d\n", &x, &y)
			if l == 0 {
				fmt.Printf("Good Bue!")
				return
			}
			//_, err := fmt.Sscanf(str, ")
			if err != nil {
				fmt.Printf("Invalid data - try again: %s\n", err.Error())
				continue
			}
			move = Move{x, y, player}
		}
		err := ttt.Play(move)
		if err != nil {
			fmt.Printf("Invalid move - try again: %s\n", err)
			continue
		}
		ttt.Print()
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
	fmt.Printf("Congratulations %s - you WIN!\n", winner)

}
