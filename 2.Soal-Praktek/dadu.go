package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	ListDadu []int64
	Point    int
}

// Generate Random Value of Dadu per Player Active
func GenerateDadu(dadu int64) []int64 {
	val := []int64{}
	count := int64(1)
	for count <= dadu {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(6)
		val = append(val, int64(n)+1)
		count++
	}
	return val
}

func GameDadu(player, dadu int64) string {

	// Initiate Variable playerActive and Variable Counter
	var playerActive []Player
	var counter []Player

	// Initiate Players Name and Length of the List Of Dadu Per Players
	for i := int64(1); i <= player; i++ {
		var gamer Player
		gamer.Name = fmt.Sprintf("Player %d", i)
		counter = append(counter, gamer)
		gamer.ListDadu = make([]int64, dadu)
		playerActive = append(playerActive, gamer)
	}

	rules := int64(0)
	winner := ""
	counterPoint := 0

	// Looping Until Player Active = 1 or Player have finished = Total Player - 1
	for rules != player-1 {
		for i, val := range playerActive {
			if len(val.ListDadu) != 0 {
				// Generate Dadu for Player Active
				playerActive[i].ListDadu = GenerateDadu(int64(len(val.ListDadu)))
				// Selection Value of Dadu
				for _, v := range playerActive[i].ListDadu {
					if v == 1 {
						if i+1 == len(playerActive) {
							counter[0].ListDadu = append(counter[0].ListDadu, 1)
						} else {
							counter[i+1].ListDadu = append(counter[i+1].ListDadu, 1)
						}
					} else if v == 6 {
						// Add Point if there is number 6 in listDadu
						playerActive[i].Point++
					} else {
						counter[i].ListDadu = append(counter[i].ListDadu, v)
					}
				}
			}
		}

		// Monitoring the Game Play and Check the Changes
		fmt.Println("playerActive", playerActive)
		fmt.Println("counter", counter)

		// Replace ListDadu In Player Active with Counter Data and Reset ListDadu of Counter
		for i, v := range counter {
			playerActive[i].ListDadu = v.ListDadu
			counter[i].ListDadu = []int64{}
		}
		// Search Player have Finished
		for i, v := range playerActive {
			if len(v.ListDadu) == 0 {
				// Increament Rules if there is player have finished
				rules++
				// Get The Winner that have Highest Point,
				// If there are more than 1 that have Highest Point, Priority the First Player have finished
				if v.Point > counterPoint {
					counterPoint = playerActive[i].Point
					winner = v.Name
				}
				// Remove the Player that have Finished from PlayerActive List
				playerActive = append(playerActive[:i], playerActive[i+1:]...)
				counter = append(counter[:i], counter[i+1:]...)
				break
			}
		}
		// Tracking The Data Active Player and The Winners After Evaluation
		fmt.Println("Check Data Active Player:", playerActive)
		fmt.Println("Winners: ", winner)
		fmt.Println()
	}
	return winner
}

func main() {
	totalPlayer := int64(20)
	totalListDadu := int64(25)
	fmt.Println("The Winner Is:", GameDadu(totalPlayer, totalListDadu))
	// Set Time For Goroutine Run Until Finish the Process
	// The Time Addapt With The Value Of Total Player and Total List Dadu Per Player
	time.Sleep(time.Millisecond)
}
