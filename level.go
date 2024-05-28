package main

import rl "github.com/gen2brain/raylib-go/raylib"

var cam = rl.Camera2D{
	Target:   rl.Vector2{X: 0, Y: 0},
	Offset:   rl.Vector2{X: 0, Y: 0},
	Rotation: 0.0,
	Zoom:     1.0,
}

func drawlevel() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(cam)

	drawPlayer()

	rl.EndMode2D()
	rl.EndDrawing()

	updatePlayer()
}
