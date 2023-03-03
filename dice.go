package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func diceGame(pemain, dadu int) int {
	// seed random generator by time now
	rand.Seed(time.Now().UnixNano())

	// roll dadu pemain
	dice := make([][]int, pemain)
	for i := range dice {
		dice[i] = make([]int, dadu)
		for j := range dice[i] {
			dice[i][j] = rand.Intn(6) + 1
		}
	}

	var playerPoin = map[int]int{}

	// Main loop permainan
	for {
		// Lepaskan dadu dari setiap pemain
		for i := range dice {
			for j := range dice[i] {
				dice[i][j] = rand.Intn(6) + 1
			}
		}

		// Evaluasi hasil lemparan dadu

		for i := range dice {

			for j := range dice[i] {
				if dice[i][j] == 6 {
					// Keluarkan dadu 6 dan tambahkan 1 poin
					dice[i][j] = 0
					numPlayer := fmt.Sprintf("%d", i+1)
					num, _ := strconv.Atoi(numPlayer)
					playerPoin[num]++
					// fmt.Printf("Pemain %d mendapat poin!\n", i+1)

				} else if dice[i][j] == 1 {
					// Berikan dadu 1 ke pemain di sebelah kanan
					if i == pemain-1 {
						dice[0] = append(dice[0], 1)
					} else {
						dice[i+1] = append(dice[i+1], 1)
					}
					dice[i][j] = 0
				}
			}
		}

		// Hapus dadu yang sudah dikeluarkan atau diberikan
		for i := range dice {
			for j := 0; j < len(dice[i]); {
				if dice[i][j] == 0 {
					dice[i] = append(dice[i][:j], dice[i][j+1:]...)
				} else {
					j++
				}
			}
		}

		// Cek apakah permainan sudah selesai
		numActivePlayers := 0
		for i := range dice {
			if len(dice[i]) > 0 {
				numActivePlayers++
			}
		}
		if numActivePlayers == 1 {
			break
		}
	}

	// Cari pemain dengan poin tertinggi
	max := 0
	winner := 0
	for i, v := range playerPoin {
		if v > max {
			max = v
			winner = i
		}
	}
	fmt.Printf("Pemain %d menang dengan %d poin!\n", winner, max)
	return winner
}

func main() {
	fmt.Println(diceGame(3, 4))
}
