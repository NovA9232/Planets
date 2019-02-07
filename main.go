package main

import (
  "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
	"strconv"

	"planet"
)

const (
  SCREEN_W float32 = 1000
  SCREEN_H float32 = 800
)

var (
	World *box2d.World
	done = make(chan bool)
	paused bool = false
)

func Draw() {
	for i := 0; i < len(planet.Planets); i++ {
		planet.Planets[i].Draw()
	}
}

func Update(dt float64) {
	World.Step(dt)
	for i := 0; i < len(planet.Planets); i++ {
		go planet.Planets[i].Update()
	}
}

func makeSqr(offsetX, offsetY, r float64, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			planet.Planets = append(planet.Planets, planet.NewPlanet(len(planet.Planets), ((float64(i)*r*2)+offsetX), ((float64(j)*r*2)+offsetY), 0, 0, r))
		}
	}
}


func checkInputs(mouseDown *bool, startMX, startMY, endMX, endMY *int32) {
	if rl.IsMouseButtonDown(0) && !(*mouseDown) {
		*mouseDown = true
		*startMX, *startMY = rl.GetMouseX(), rl.GetMouseY()
	}

	if rl.IsMouseButtonReleased(0) {
		*mouseDown = false
		*endMX, *endMY = rl.GetMouseX(), rl.GetMouseY()
		println(*endMX-*startMX, *endMY-*startMY)
		if rl.IsMouseButtonReleased(0) {
			planet.Planets = append(planet.Planets, planet.NewPlanet(len(planet.Planets), float64(*startMX), float64(*startMY), float64(*startMX-*endMX), float64(*startMY-*endMY), 5))
		}
	}

	if rl.IsKeyPressed(rl.KeyP) {
		paused = !paused
	}
	if rl.IsKeyPressed(rl.KeyR) {
		reset()
	}
}

func reset() {
	World.Clear()
	planet.Planets = []*planet.Planet{}
}

func main() {
  rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Planets")
  rl.SetTargetFPS(60)

	var (
		mouseDown bool  = false
		startMX   int32 = 0
		startMY   int32 = 0
		endMX     int32 = 0
		endMY     int32 = 0
	)

	planet.Planets = []*planet.Planet{}
	World = box2d.NewWorld(box2d.Vec2{0, 0}, 20)
	planet.World = World

	//planet.Planets = append(planet.Planets, planet.NewPlanet(int32(len(planet.Planets)), 300, 300, 0, 0, 5))
	//planet.Planets = append(planet.Planets, planet.NewPlanet(int32(len(planet.Planets)), 300, 400, 0, 0, 5))
	makeSqr(200, 200, 1, 20, 2)

  for !rl.WindowShouldClose() {
		if !paused {
			Update(float64(rl.GetFrameTime()))
		}
		checkInputs(&mouseDown, &startMX, &startMY, &endMX, &endMY)

    rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		Draw()

		rl.DrawFPS(10, 10)
		rl.DrawText("Body count: "+strconv.Itoa(len(planet.Planets)), 10, 40, 12, rl.RayWhite)

		rl.EndDrawing()
  }

  rl.CloseWindow()
}
