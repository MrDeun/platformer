package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	position rl.Vector2
	width    uint
	height   uint
	color    rl.Color
}

type Ball struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
	color    rl.Color
}

func CreateBall(position rl.Vector2, velocity rl.Vector2, radius float32, color rl.Color) Ball {
	return Ball{position: position, velocity: velocity, radius: radius, color: color}
}

func CreatePaddle(position rl.Vector2, width uint, height uint, color rl.Color) Paddle {
	return Paddle{position: position, width: width, height: height, color: color}
}

func (_paddle *Paddle) draw() {
	rl.DrawRectangle(
		int32(_paddle.position.X)-int32(_paddle.width)/2,
		int32(_paddle.position.Y)/int32(_paddle.height),
		int32(_paddle.width),
		int32(_paddle.height),
		_paddle.color)
}

func (_ball *Ball) draw() {
	rl.DrawCircle(int32(_ball.position.X)-int32(_ball.radius)/2, int32(_ball.position.Y)-int32(_ball.radius)/2, _ball.radius, _ball.color)
}
