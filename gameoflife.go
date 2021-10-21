package main

import (
	"gameoflife/models"

	log "github.com/sirupsen/logrus"
)

/* Game of Life Design
1. Represent the initial seed universe in a NxN array of Cells - use true/false to represent alive dead.
1.1 Since we are trying to represent infinite array, we will need N+1xN+1 rows to cover all the MxN cells
2. Each element of the array is a *Cell*
3. Each Cell can be in either of 2 states: Alive or Dead
4. For cell at index (r,c), the neighbors will be at :
		1. r-1, c-1
		2. r-1,c
		3. r-1, c+1
		4. r, c-1
		5. r, c+1
		6. r+1, c-1
		7. r+1, c
		8. r+1, c+1
5. Methods for each cell:
isAlive() returning bool - true if cell at (r,c) is alive, false if not - returns Cell.isAlive()
killCell() returning bool - sets isAlive to false
liveCell() - sets isAlive to true

Rules for tick:
6. Number of live neighbors for a cell at (r,c) => func CountLiveNeighbors(r,c) int
7. If isAlive(r,c) && If CountLiveNeighbors(r,c) < 2:
	KillCell(r,c) - due to underpopulation
8. if isAlive(r,c) && CountLiveNeighbors(r,c) >=2 and <=3 : no change
9. If isAlive(r,c) && CountLiveNeighbors(r,c) > 3:
	KillCell(r,c) - due to overpopulation
10. !isAlive(r,c) && CountLiveNeighbors(r,c) == 3: liveCell(r,c)
11. Steps 7 - 10 must happen simultaneously i.e. Events are r-1,c cannot affect events for r,c?
12. We need a method to kick off the tick and to simulate the display the state of the universe doPrintUniverse() at each step.
13. End State: When all the cells are dead? Everything that has a beginning has an end?
14. Can there be a "quantum-stable state" where the state of the universe never changes? Another possible end state?
15. We will need to keep track of state of the universe across ticks.
*/

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
