package planet

import (
	"github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

const (
	//G float64 = 0.0000000000667
	G float64 = 0.000667
	PL_DENSITY float64 = 5514
	PL_FRICTION float64 = 0
)

var (
	World *box2d.World
	Planets []*Planet
)

type Planet struct {
	box2d.Body
	ID int
	Rad float64
}

func getMass(d, r float64) float64 {
	area := rl.Pi*(r*r)   // Here area is volume as this is 2d
	return d*area
}

func NewPlanet(id int, x, y, vx, vy, r float64) *Planet {
	p := Planet{
		ID: id,
		Rad: r,
	}
	p.Set(&box2d.Vec2{r*2, r*2}, getMass(PL_DENSITY, r))
	p.Body.Friction = PL_FRICTION
	p.Body.Position = box2d.Vec2{x, y}
	p.Body.Velocity = box2d.Vec2{vx, vy}
	World.AddBody(&p.Body)
	return &p
}

func (self *Planet) getGravForce() box2d.Vec2 {
	var fTotal = box2d.Vec2{0, 0}
	for i := 0; i < len(Planets); i++ {
		if Planets[i].ID != self.ID {
			distVec, dist := GetDistance(self.Body.Position, Planets[i].Position)
			f := G * (float64(self.Body.Mass*Planets[i].Body.Mass) / (dist * dist))
			fTotal.X += f * distVec.X
			fTotal.Y += f * distVec.Y
		}
	}
	return fTotal
}

func (self *Planet) Update() {
	self.Body.Force = self.getGravForce()
}

func (self *Planet) Draw() {
	rl.DrawCircleLines(int32(self.Body.Position.X), int32(self.Body.Position.Y), float32(self.Rad), rl.RayWhite)
}
