// storing definitions

package main

import rl "github.com/gen2brain/raylib-go/raylib"

const GRID_COLS int = 20
const GRID_ROWS int = 25

type CellCode int
type GameState [GRID_ROWS][GRID_COLS]CellCode

const (
	EMPTY CellCode = iota
	WALL
	PACMAN
	FOOD
	SPECIAL_FOOD
	GHOST
)

// colors
var (
	NavyBlue = rl.Color{R: 20, G: 14, B: 57, A: 0}
	Blue     = rl.Blue
)
