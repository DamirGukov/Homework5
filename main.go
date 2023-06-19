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

func FieldForPlaying(field [3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("{ %s } { %s } { %s }\n", field[i][0], field[i][1], field[i][2])

	}
}

func CheckWin(field [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if field[0][0] == field[0][1] && field[0][0] == field[0][2] && field[0][0] != empty && field[0][1] != empty && field[0][2] != empty {
			return true
		} else if field[1][0] == field[1][1] && field[1][0] == field[1][2] && field[1][0] != empty && field[1][1] != empty && field[1][2] != empty {
			return true
		} else if field[2][0] == field[2][1] && field[2][0] == field[2][2] && field[2][0] != empty && field[2][1] != empty && field[2][2] != empty {
			return true
		} else if field[0][0] == field[1][0] && field[0][0] == field[2][0] && field[0][0] != empty && field[1][0] != empty && field[2][0] != empty {
			return true
		} else if field[0][1] == field[1][1] && field[2][1] == field[0][1] && field[0][1] != empty && field[1][1] != empty && field[2][1] != empty {
			return true
		} else if field[0][2] == field[1][2] && field[2][2] == field[0][2] && field[0][2] != empty && field[1][2] != empty && field[2][2] != empty {
			return true
		}

		if field[0][0] == field[1][1] && field[0][0] == field[2][2] && field[0][0] != empty && field[1][1] != empty && field[2][2] != empty {
			return true
		} else if field[2][0] == field[1][1] && field[2][0] == field[0][2] && field[2][0] != empty && field[2][1] != empty && field[0][2] != empty {
			return true
		}

	}
	return false
}

func CheckDraw(field [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field[i][j] == empty {
				return false
			}
		}
	}
	return true
}

func ClearField(field *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			field[i][j] = empty
		}
	}
}

func MakeStep(field *[3][3]string, randomPlayer int, scores map[int]int) map[int]int {
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

		if field[row][col] != empty {
			fmt.Println("This cell is already occupied. Choose another one!")
			continue
		}

		fmt.Println("You chose:", row, col)
		if randomPlayer == player1 {
			field[row][col] = X
		} else if randomPlayer == player2 {
			field[row][col] = O
		}

		FieldForPlaying(*field)

		if randomPlayer == player1 {
			randomPlayer = player2
		} else {
			randomPlayer = player1
		}

		if CheckWin(*field) {
			fmt.Println("Player", randomPlayer, "wins!")
			if player1 == randomPlayer {
				scores[1]++
			} else if player2 == randomPlayer {
				scores[2]++
			}
			break
		} else if CheckDraw(*field) {
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

	field := [3][3]string{}
	scores := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}

	for {
		var newGame string
		randomPlayer := rand.Intn(2) + 1

		FieldForPlaying(field)
		scores = MakeStep(&field, randomPlayer, scores)

		fmt.Println("Player1 wins:", scores[1])
		fmt.Println("Player2 wins:", scores[2])
		fmt.Println("Draw:", scores[3])

		fmt.Println("Do you want to start a new game? yes/no")
		fmt.Scanln(&newGame)

		if newGame != "yes" {
			break
		}

		ClearField(&field)
		FieldForPlaying(field)
	}
}
