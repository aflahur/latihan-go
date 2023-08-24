package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numPlayers := getInput("Masukkan jumlah pemain: ")
	numDice := getInput("Masukkan jumlah dadu: ")

	dice := make([][]int, numPlayers)
	for i := range dice {
		dice[i] = make([]int, numDice)
		for j := range dice[i] {
			dice[i][j] = rollDice()
		}
	}

	playGame(dice)
}

func getInput(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func rollDice() int {
	return rand.Intn(6) + 1
}

func removeZeros(slice []int) []int {
	result := []int{}
	for _, val := range slice {
		if val != 0 {
			result = append(result, val)
		}
	}
	return result
}

func findRemainingPlayer(dice [][]int) int {
	for i, playerDice := range dice {
		if len(playerDice) > 0 {
			return i
		}
	}
	return -1
}

func findWinner(dice [][]int) int {
	maxPoints := 0
	winner := -1
	for i, playerDice := range dice {
		points := 0
		for _, val := range playerDice {
			points += val
		}
		if points > maxPoints {
			maxPoints = points
			winner = i
		}
	}
	return winner
}

func playGame(dice [][]int) {
	round := 0
	for {
		round++
		fmt.Printf("==================\n")
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		remainingPlayers := 0
		for i, playerDice := range dice {
			if len(playerDice) > 0 {
				remainingPlayers++
				fmt.Printf("Pemain #%d (%d): %v\n", i+1, len(playerDice), playerDice)
			}
		}

		if remainingPlayers <= 1 {
			break
		}

		for i, playerDice := range dice {
			if len(playerDice) == 0 {
				continue
			}

			for j := range playerDice {
				if playerDice[j] == 1 {
					if i < len(dice)-1 {
						dice[i+1] = append(dice[i+1], 1)
					} else {
						dice[0] = append(dice[0], 1)
					}
					playerDice[j] = 0
				} else if playerDice[j] == 6 {
					playerDice[j] = 0
				}
			}

			dice[i] = removeZeros(playerDice)
		}

		fmt.Printf("Setelah evaluasi:\n")
		for i, playerDice := range dice {
			if len(playerDice) > 0 {
				fmt.Printf("Pemain #%d (%d): %v\n", i+1, len(playerDice), playerDice)
			}
		}
	}

	fmt.Printf("==================\n")
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", findRemainingPlayer(dice)+1)
	winner := findWinner(dice)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", winner+1)
}
