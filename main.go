package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const windowX = 1600
const windowY = 900

const centerX = windowX / 2
const centerY = windowY / 2

var LightBlue = rl.NewColor(164, 221, 237, 255)
var LightRed = rl.NewColor(236, 83, 83, 0)

func main() {
	rl.InitWindow(windowX, windowY, "Hey")
	defer rl.CloseWindow()

	_ball := CreateBall(rl.Vector2{X: centerX, Y: centerY}, rl.Vector2Zero(), 8.0, rl.White)
	_paddle1 := CreatePaddle(rl.Vector2{X: centerX + 400, Y: 0}, 16, 80, LightRed)
	_paddle2 := CreatePaddle(rl.Vector2{X: centerX - 400, Y: 0}, 16, 80, LightRed)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		_ball.draw()
		_paddle1.draw()
		_paddle2.draw()
		rl.DrawText("Hello world!", centerX, centerY, 16.0, rl.White)
		rl.EndDrawing()
	}
}
