package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	phy "github.com/gen2brain/raylib-go/physics"
	"math"

	"bodies"
)

const (
	WINDOW_W = 1200
	WINDOW_H = 800

	PL_DENSITY = 5
)

var (
	G float32 = float32(6.67*(math.Pow(10, -8))) // Not a constant unfortunately
)

func main() {
	bodies.G = G

	rl.InitWindow(WINDOW_W, WINDOW_H, "Test")
	rl.SetTargetFPS(60)

	phy.Init()
	phy.SetGravity(0, 0)
	phy.SetPhysicsTimeStep(1/60)

	// bodies.NewPlanet(len(bodies.ListOfBods), 100, 100, 50, PL_DENSITY)
	// bodies.NewPlanet(len(bodies.ListOfBods), 100, 220, 50, PL_DENSITY)
	// bodies.NewPlanet(len(bodies.ListOfBods), 300, 150, 50, PL_DENSITY)

	for i := 0; i < 50; i++ {
		bodies.NewPlanet(len(bodies.ListOfBods), i*20, 300, 10, PL_DENSITY)
	}

	for !rl.WindowShouldClose() {
		bodies.Update()
		phy.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		bodies.DrawBodies()

		rl.EndDrawing()

		phy.SetPhysicsTimeStep(rl.GetFrameTime())
	}

	phy.Close()

	rl.CloseWindow()
}
