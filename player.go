package main

import rl "github.com/gen2brain/raylib-go/raylib"

var posx float32 = 0
var posy float32 = 0
var pxx, pxy = tile2pixel(posx, posy)
var wd, hd = tile2pixel(3, 3)

var speed float32 = 13

func drawPlayer() {
	rl.DrawRectangle(pxx, pxy, wd, hd, rl.Red)
}

func updatePlayer() {
	if rl.IsKeyDown(rl.KeyW) {
		posy -= speed * rl.GetFrameTime()
	} else if rl.IsKeyDown(rl.KeyS) {
		posy += speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyA) {
		posx -= speed * rl.GetFrameTime()
	} else if rl.IsKeyDown(rl.KeyD) {
		posx += speed * rl.GetFrameTime()
	}

	pxx, pxy = tile2pixel(posx, posy)
}
