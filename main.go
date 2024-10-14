package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	game "github.com/mrdeun/platformer/internals"
)

const windowX = 1600
const windowY = 900

const centerX = windowX / 2
const centerY = windowY / 2

func main() {
	rl.InitWindow(windowX, windowY, "Hey")
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Hello world!", centerX, centerY, 16.0, rl.White)
		rl.EndDrawing()
	}
}
