package models

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Universe represents the state of the Universe at any instant
type Universe struct {
	// rows the number of rows that represent a snap-shot of the infinite universe
	rows int
	// cols the number of columns that represent a snap-shot of the infinite universe
	cols int
	// previous holds the previous state of the universe
	previous [][]Cell
	// current holds the current state which will result from applying rules on previous state
	current [][]Cell
}

// init2DCellArray initializes the initial state of the universe
func init2DCellArray(r, c int) [][]Cell {
	var arr [][]Cell
	if r > 0 && c > 0 {
		arr = make([][]Cell, r)
		for i := range arr {
			arr[i] = make([]Cell, c)
		}
	}
	return arr
}

// GetCellFromPrev returns an existing cell in the universe's state else a dead cell if outside the current scope of the universe state
func (u *Universe) GetCellFromPrev(r, c int) Cell {
	if r < 0 || r > u.rows-1 || c < 0 || c > u.cols-1 {
		return NewCell(0, 0)
	}
	return u.previous[r][c]
}

// InitUniverse initializes the state of the universe from a seed state
func InitUniverse(seed [][]int) (*Universe, error) {
	if len(seed) == 0 {
		return nil, errors.New("empty/invalid seed array")
	}

	var u Universe
	u.rows = len(seed) + 2
	u.cols = len(seed[0]) + 2

	u.previous = init2DCellArray(u.rows, u.cols)
	u.current = nil

	for i := 0; i < u.rows-2; i++ {
		for j := 0; j < u.cols-2; j++ {
			u.previous[i+1][j+1] = NewCell(seed[i][j], 0)
		}
	}

	return &u, nil
}

// String produces a string representation of current state of the universe
func (u Universe) String() string {
	op := ""
	for i := 0; i < u.rows; i++ {
		for j := 0; j < u.cols-1; j++ {
			op += fmt.Sprintf("%s ", u.previous[i][j].String())
		}
		op += fmt.Sprintf("%s\n", u.previous[i][u.cols-1].String())
	}
	return op
}

// IsStateChanged compares the current state of the universe to the previous state
// returns boolean true if state has changed. false otherwise
func (u *Universe) IsStateChanged() bool {
	if u.current == nil {
		return true
	}

	for i := 0; i < u.rows; i++ {
		for j := 0; j < u.cols; j++ {
			if u.current[i][j].alive != u.previous[i][j].alive {
				return true
			}
		}
	}
	return false
}

// Tick triggers the change in state for the entire universe
// Current state of all cells are updated based on their previous state
// All changes happen simultaneously.
// Returns true - if the state of the universe has changed
// False - if no change was detected
func (u *Universe) Tick() bool {
	u.current = init2DCellArray(u.rows, u.cols)

	// update state from previous
	for i := 0; i < u.rows; i++ {
		for j := 0; j < u.cols; j++ {
			cc := &Cell{
				alive: u.GetCellFromPrev(i, j).alive,
			}
			cc.CountLiveNeighbors([]Cell{
				u.GetCellFromPrev(i-1, j-1),
				u.GetCellFromPrev(i-1, j),
				u.GetCellFromPrev(i-1, j+1),
				u.GetCellFromPrev(i, j-1),
				u.GetCellFromPrev(i, j+1),
				u.GetCellFromPrev(i+1, j-1),
				u.GetCellFromPrev(i+1, j),
				u.GetCellFromPrev(i+1, j+1),
			})
			u.current[i][j] = *cc
		}
	}

	// update next generation state
	for i := 1; i < u.rows-1; i++ {
		for j := 1; j < u.cols-1; j++ {
			cc := u.current[i][j]
			const (
				lowerLimitPopulation = 2
				upperLimitPopulation = 3
			)

			// apply game rules
			if cc.IsAlive() {
				if cc.liveNeighbors < lowerLimitPopulation {
					log.Infof("Apply lower limit population rule on cell (%d,%d)", i, j)
					cc.KillCell()
				} else if cc.liveNeighbors > upperLimitPopulation {
					log.Infof("Apply upper limit population rule on cell (%d,%d)", i, j)
					cc.KillCell()
				}
			} else {
				if cc.liveNeighbors == upperLimitPopulation {
					log.Infof("Apply alive rule on cell (%d,%d)", i, j)
					cc.LiveCell()
				}
			}

			u.current[i][j] = cc
		}
	}

	if !u.IsStateChanged() {
		return false
	}

	u.previous = u.current
	return true
}

// StartGame triggers the state changes in universe.
// Returns when the Tick() stops producing any change of state in universe
func (u *Universe) StartGame() error {
	if u == nil {
		return errors.New("universe state unknown")
	}

	fmt.Println(u)

	for u.Tick() {
		fmt.Println(u)
	}
	return nil
}
