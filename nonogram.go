package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	var fontSize int32 = int32(math.Floor(screenHeight / 20))

	rl.InitWindow(screenWidth, screenHeight, "Nonograms made with Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!",
			(screenWidth-rl.MeasureText("Congrats! You created your first window!", fontSize))/2,
			(screenHeight-fontSize)/2,
			fontSize,
			rl.LightGray)

		rl.EndDrawing()
	}
}
