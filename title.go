package main

import rl "github.com/gen2brain/raylib-go/raylib"

var inmenu = false

func drawtMenu() {
	// TODO: This is horrible please rewite it.
	if !inmenu {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		mousePoint := rl.GetMousePosition()

		// Title
		tposx, tposy := tile2pixel((tiles.X / 2), 3)
		tposx -= fitscreen(40)
		rl.DrawText("Gmae?", tposx, tposy, fitscreen(40), rl.LightGray)

		// Start button
		c_x, p_y := tile2pixel(tiles.X/2, tiles.Y/2)
		p_x := c_x - fitscreen(200)
		c_x -= fitscreen(40)
		c_y := p_y + fitscreen(20)

		bound := rl.NewRectangle(float32(p_x), float32(p_y), float32(fitscreen(400)), float32(fitscreen(80)))
		if rl.CheckCollisionPointRec(mousePoint, bound) {
			rl.DrawRectangle(p_x, p_y, fitscreen(400), fitscreen(80), rl.Green)
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				rl.EndDrawing()
				inmenu = true
			}
		} else {
			rl.DrawRectangle(p_x, p_y, fitscreen(400), fitscreen(80), rl.Red)
		}

		rl.DrawText("PLAY", c_x, c_y, fitscreen(40), rl.White)

		rl.EndDrawing()
	} else {
		drawlevel()
	}
}
