package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_HEIGHT int32 = 800
const WINDOW_WIDTH int32 = 800
const FPS int32 = 20

var is_paused bool = false
var wall_map WallMap

func initializeGameState() {
	// to be implemented
}

func handleKeyPress(key int32) {
	switch key {
	case rl.KeyP:
		is_paused = !is_paused
		break
	default:
		break
	}
}

func drawHeaderArea() {
	rl.DrawText("SCORE", WINDOW_WIDTH / 2, 20, 16, rl.White)
}

func drawBottomArea() {
	// to be implemented
}

func drawCell(row_pos int, col_pos int) {
	// to be implemented
}

func drawWalls() {
	for i, row := range wall_map {
		for j, val := range row {
			if val {
				drawCell(i, j)
			}
		}
	}
}

func handleDrawing() {
	rl.BeginDrawing()
	rl.ClearBackground(NavyBlue)

	drawHeaderArea()
	drawBottomArea()
	rl.EndDrawing()
}

func main() {
	initializeGameState()

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Pac-Man")
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)
	rl.SetExitKey(0) // not to use ESC key to exit

	drawWalls()

	for !rl.WindowShouldClose() {
		var key int32 = rl.GetKeyPressed()
		if key != 0 {
			handleKeyPress(key)
		}
		if !is_paused {
			handleDrawing()
		}
	}
}
