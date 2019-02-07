package main

import (
  "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"planet"
)

const (
  SCREEN_W float32 = 1000
  SCREEN_H float32 = 800
)

var (
	World *box2d.World
	Planets []*planet.Planet
)

func Update(dt float64) {
	World.Step(dt)
	for i := 0; i < len(Planets); i++ {
		Planets[i].Update()
	}
}

func main() {
  rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Planets")
  rl.SetTargetFPS(60)

	World = box2d.NewWorld(box2d.Vec2{0, 0}, 10)
	planet.World = World

  for !rl.WindowShouldClose() {
		Update(float64(rl.GetFrameTime()))

    rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
  }

  rl.CloseWindow()
}
