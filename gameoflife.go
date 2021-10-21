package main

import (
	"gameoflife/models"

	log "github.com/sirupsen/logrus"
)

func main() {
	// seed array from example
	seed := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	u, err := models.InitUniverse(seed)
	if err != nil {
		log.Fatal(err)
	}
	u.StartGame()
}
