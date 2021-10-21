package models

import (
	"errors"
)

// Cell is individual unit of life in the Universe
type Cell struct {
	// alive is set to true if the cell is alive
	alive         bool
	// liveNeighbors keeps track of the number of adjacent neighbors
	// where alive is true for this cell
	liveNeighbors int
}

// String produces string representation of the cell's state
// returns: 'A' if cell.alive is true, 
// 'D' if false
func (c Cell) String() string {
	if c.alive {
		return "A"
	}
	return "D"
}

// IsAlive is a getter function for the live state of the cell
func (c Cell) IsAlive() bool {
	return c.alive
}

// NewCell returns a cell struct based on user input 
func NewCell(a, nc int) Cell {
	if a == 1 {
		return Cell{alive: true}
	}
	return Cell{alive: false}
}

// KillCell sets the live status to false
func (c *Cell) KillCell() {
	c.alive = false
}

// KillCell sets the live status to true
func (c *Cell) LiveCell() {
	c.alive = true
}

// CountLiveNeighbors sets the count of cell's adjacent based on array of surrounding cells
func (c *Cell) CountLiveNeighbors(neighbors []Cell) error {
	if neighbors == nil || len(neighbors) != 8 {
		return errors.New("invalid neighbors list")
	}

	c.liveNeighbors = 0
	// count live neighbors
	for _, neighborCell := range neighbors {
		if neighborCell.alive {
			c.liveNeighbors++
		}
	}
	return nil
}
