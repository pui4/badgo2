package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Tile functions
const TILESIZE int = 16

var tiles rl.Vector2

func tile2pixel(tilex float32, tiley float32) (int32, int32) {
	xp := int32(tilex * float32(TILESIZE))
	yp := int32(tiley * float32(TILESIZE))
	return xp, yp
}

func fitscreen(size float32) int32 {
	x := size / float32(720)
	return int32(x * float32(rl.GetScreenHeight()))
}
