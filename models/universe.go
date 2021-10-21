package models

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Universe struct {
	rows     int
	cols     int
	previous [][]Cell
	current  [][]Cell
}

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

func (u *Universe) Tick() bool {
	u.current = init2DCellArray(u.rows, u.cols)
	// update state from previous
	for i := 1; i < u.rows-1; i++ {
		for j := 1; j < u.cols-1; j++ {
			cc := &Cell{
				alive: u.previous[i][j].alive,
			}
			cc.CountLiveNeighbors([]Cell{
				u.previous[i-1][j-1],
				u.previous[i-1][j],
				u.previous[i-1][j+1],
				u.previous[i][j-1],
				u.previous[i][j+1],
				u.previous[i+1][j-1],
				u.previous[i+1][j],
				u.previous[i+1][j+1],
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
					log.Debugf("Apply lower limit population rule on cell (%d,%d)", i, j)
					cc.KillCell()
				} else if cc.liveNeighbors > upperLimitPopulation {
					log.Debugf("Apply upper limit population rule on cell (%d,%d)", i, j)
					cc.KillCell()
				}
			} else {
				if cc.liveNeighbors == upperLimitPopulation {
					log.Debugf("Apply alive rule on cell (%d,%d)", i, j)
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
