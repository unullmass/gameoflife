package models

import (
	"errors"
)

type Cell struct {
	alive         bool
	liveNeighbors int
}

func (c Cell) String() string {
	if c.alive {
		return "A"
	}
	return "D"
}

func (c Cell) IsAlive() bool {
	return c.alive
}

func NewCell(a, nc int) Cell {
	if a == 1 {
		return Cell{alive: true}
	}
	return Cell{alive: false}
}

func (c *Cell) KillCell() {
	c.alive = false
}

func (c *Cell) LiveCell() {
	c.alive = true
}

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
