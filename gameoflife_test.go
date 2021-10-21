package main

import (
	"gameoflife/models"
	"testing"
)

func TestExample1(t *testing.T) {
	// seed array from example
	seed1 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	u, err := models.InitUniverse(seed1)
	if err != nil {
		t.Fatal(err)
	}
	u.StartGame()

}

func TestExample2(t *testing.T) {
	seed := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
	}
	u, err := models.InitUniverse(seed)
	if err != nil {
		t.Fatal(err)
	}
	u.StartGame()
}

func TestAllAlive(t *testing.T) {
	seed := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	u, err := models.InitUniverse(seed)
	if err != nil {
		t.Fatal(err)
	}
	u.StartGame()
}

func TestAllDead(t *testing.T) {
	seed := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	u, err := models.InitUniverse(seed)
	if err != nil {
		t.Fatal(err)
	}
	u.StartGame()
}

func TestInvalidSeed(t *testing.T) {
	seed := [][]int{}
	_, err := models.InitUniverse(seed)
	if err == nil {
		t.Fail()
	}
}

func TestSmall(t *testing.T) {
	seed := [][]int{
		{1, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
	}
	u, err := models.InitUniverse(seed)
	if err != nil {
		t.Fatal(err)
	}
	u.StartGame()
}
