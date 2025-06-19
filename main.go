package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
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
	// to be implemented
}

func drawBottomArea() {
	// to be implemented
}

func drawCell(row_pos int, col_pos int) {
	// to be implemented
}

func handleDrawing() {
	rl.BeginDrawing()
	rl.ClearBackground(NavyBlue)

	drawHeaderArea()
	for i, row := range wall_map {
		for j, val := range row {
			if val {
				drawCell(i, j)
			}
		}
	}
	drawBottomArea()
	rl.EndDrawing()
}

func main() {
	fmt.Println("Testing")
	initializeGameState()

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Pac-Man")
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)
	rl.SetExitKey(0) // not to use ESC key to exit

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
