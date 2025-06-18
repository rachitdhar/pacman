package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "Pac-Man")
	defer rl.CloseWindow()

	rl.SetTargetFPS(100)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(NavyBlue)

		rl.EndDrawing()
	}
}
