package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	phy "github.com/gen2brain/raylib-go/physics"
	//"math"

	"bodies"
)

const (
	WINDOW_W = 1200
	WINDOW_H = 800

	PL_DENSITY = 5514
)

var (
	G float32 = float32(200000000000000)//*(math.Pow(10, 1))) // Not a constant unfortunately
)

func main() {
	bodies.G = G

	rl.InitWindow(WINDOW_W, WINDOW_H, "Test")
	rl.SetTargetFPS(60)

	phy.Init()
	phy.SetGravity(0, 0)
	phy.SetPhysicsTimeStep(1/60)
	pl := new(bodies.Planet)
	bodies.ListOfBods = append(bodies.ListOfBods, pl)
	pl.Init(len(bodies.ListOfBods), 100, 100, 50, PL_DENSITY)
	pl.Init(len(bodies.ListOfBods), 100, 220, 50, PL_DENSITY)

	for !rl.WindowShouldClose() {
		phy.Update()

		bodies.UpdateGravity()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		bodies.DrawBodies()

		rl.EndDrawing()

		phy.SetPhysicsTimeStep(rl.GetFrameTime())
	}

	phy.Close()

	rl.CloseWindow()
}
