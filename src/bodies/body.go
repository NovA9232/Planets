package bodies

import (
	phy "github.com/gen2brain/raylib-go/physics"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	//"fmt"
)

var (
	G float32 // Will be set by main
	ListOfBods []Body
)

// This file contains functions for all bodies.

type Body interface { // Body interface to work with multiple bodies.
	getBody() *phy.Body
	SetForce(fx, fy float32)
}

func DrawBodies() {
	for i, body := range phy.GetBodies() {
		vertexCount := phy.GetShapeVerticesCount(i)
		for j := 0; j < vertexCount; j++ {
			vertexA := body.GetShapeVertex(j)
			jj := 0
			if j+1 < vertexCount {
				jj = j + 1
			}
			vertexB := body.GetShapeVertex(jj)
			rl.DrawLineV(vertexA, vertexB, rl.RayWhite)
		}
	}
}

func Update() { // Update function for all bodies
	UpdateGravity()
}

func UpdateGravity() {
	done := make(chan int)
	for _, bod := range ListOfBods {
		go getGravity(bod, done)
	}
	for i := 0; i < len(ListOfBods); i++ {
		<-done // Wait for go routine's to be done
	}
	close(done)
}

func getGravity(from Body, done chan<- int) {
	var fx, fy float32
	fromBody := from.getBody()
	for _, other := range phy.GetBodies() {
		if fromBody != other {
			distX, distY := other.Position.X - fromBody.Position.X, other.Position.Y - fromBody.Position.Y
			dist := math.Sqrt(float64(distX*distX + distY*distY))
			f := (G * fromBody.Mass * other.Mass)/(float32(math.Pow(dist, 2)))
			fx += f*distX
			fy += f*distY
		}
	}
	//fmt.Println(fx, fy, "force.")
	from.SetForce(fx, fy)
	done<- 1
	//from.setForce(fx, fy)
}
