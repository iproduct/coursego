package main

import (
	"fmt"
	"strings"
)

type Move struct {
	x, y   int
	player string
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

func (t *TicTacToe) Finished() (bool, string) {
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

func (t *TicTacToe) FinishedOld() bool {
	for i, row := range *t {
		for j := range row {
			if (*t)[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

var ttt TicTacToe

func main() {
	ttt.Init(3)
	ttt.Print()
	player := "X"
	var winner string
	for finished := false; !finished; finished, winner = ttt.Finished() {
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
			fmt.Printf("Invalid data: %s\n", err)
			continue
		}
		err = ttt.Play(Move{x, y, player})
		if err != nil {
			fmt.Printf("Invalid move - try again: %s\n", err)
			fmt.Errorf("Invalid move - try again: %w", err)
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
