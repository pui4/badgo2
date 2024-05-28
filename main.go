package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1280, 720, "gmae?")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update tiles
		tiles.X = float32(rl.GetScreenWidth()) / float32(TILESIZE)
		tiles.Y = float32(rl.GetScreenHeight()) / float32(TILESIZE)

		// Draw menu
		drawtMenu()
	}
}
