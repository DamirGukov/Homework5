package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	player1 = 1
	player2 = 2
	X       = "X"
	O       = "0"
	empty   = ""
)

type GameSession struct {
	field [3][3]string
}

func NewGameSession() *GameSession {
	return &GameSession{field: [3][3]string{}}

}
func (gs *GameSession) FieldForPlaying() {
	for i := 0; i < 3; i++ {
		fmt.Printf("{ %s } { %s } { %s }\n", gs.field[i][0], gs.field[i][1], gs.field[i][2])

	}
}

func (gs *GameSession) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if gs.field[0][0] == gs.field[0][1] && gs.field[0][0] == gs.field[0][2] && gs.field[0][0] != empty && gs.field[0][1] != empty && gs.field[0][2] != empty {
			return true
		} else if gs.field[1][0] == gs.field[1][1] && gs.field[1][0] == gs.field[1][2] && gs.field[1][0] != empty && gs.field[1][1] != empty && gs.field[1][2] != empty {
			return true
		} else if gs.field[2][0] == gs.field[2][1] && gs.field[2][0] == gs.field[2][2] && gs.field[2][0] != empty && gs.field[2][1] != empty && gs.field[2][2] != empty {
			return true
		} else if gs.field[0][0] == gs.field[1][0] && gs.field[0][0] == gs.field[2][0] && gs.field[0][0] != empty && gs.field[1][0] != empty && gs.field[2][0] != empty {
			return true
		} else if gs.field[0][1] == gs.field[1][1] && gs.field[2][1] == gs.field[0][1] && gs.field[0][1] != empty && gs.field[1][1] != empty && gs.field[2][1] != empty {
			return true
		} else if gs.field[0][2] == gs.field[1][2] && gs.field[2][2] == gs.field[0][2] && gs.field[0][2] != empty && gs.field[1][2] != empty && gs.field[2][2] != empty {
			return true
		}

		if gs.field[0][0] == gs.field[1][1] && gs.field[0][0] == gs.field[2][2] && gs.field[0][0] != empty && gs.field[1][1] != empty && gs.field[2][2] != empty {
			return true
		} else if gs.field[2][0] == gs.field[1][1] && gs.field[2][0] == gs.field[0][2] && gs.field[2][0] != empty && gs.field[2][1] != empty && gs.field[0][2] != empty {
			return true
		}

	}
	return false
}

func (gs *GameSession) CheckDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gs.field[i][j] == empty {
				return false
			}
		}
	}
	return true
}

func (gs *GameSession) ClearField() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			gs.field[i][j] = empty
		}
	}
}

func (gs *GameSession) MakeStep(randomPlayer int, scores map[int]int) map[int]int {
	var row, col int

	for {
		fmt.Print("Player", randomPlayer, " Enter a row in the game field (0-2): ")
		fmt.Scanln(&row)

		fmt.Print("Player", randomPlayer, " Enter a column in the game field (0-2): ")
		fmt.Scanln(&col)

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Wrong coordinates. Try again!")
			continue
		}

		if gs.field[row][col] != empty {
			fmt.Println("This cell is already occupied. Choose another one!")
			continue
		}

		fmt.Println("You chose:", row, col)
		if randomPlayer == player1 {
			gs.field[row][col] = X
		} else if randomPlayer == player2 {
			gs.field[row][col] = O
		}

		gs.FieldForPlaying()

		if randomPlayer == player1 {
			randomPlayer = player2
		} else {
			randomPlayer = player1
		}

		if gs.CheckWin() {
			fmt.Println("Player", randomPlayer, "wins!")
			if player1 == randomPlayer {
				scores[1]++
			} else if player2 == randomPlayer {
				scores[2]++
			}
			break
		} else if gs.CheckDraw() {
			fmt.Println("Draw!")
			scores[3]++
			break
		}
	}

	return scores
}

func main() {

	fmt.Println("TIC-TAC-TOE \n")
	rand.Seed(time.Now().UnixNano())

	gs := NewGameSession()
	scores := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}

	for {
		var newGame string
		randomPlayer := rand.Intn(2) + 1

		gs.FieldForPlaying()
		gs.MakeStep(randomPlayer, scores)

		fmt.Println("Player1 wins:", scores[1])
		fmt.Println("Player2 wins:", scores[2])
		fmt.Println("Draw:", scores[3])

		fmt.Println("Do you want to start a new game? yes/no")
		fmt.Scanln(&newGame)

		if newGame != "yes" {
			break
		}

		gs.ClearField()
		gs.FieldForPlaying()
	}
}
