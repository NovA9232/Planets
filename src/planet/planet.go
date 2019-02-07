package planet

import (
	"github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

var (
	World *box2d.World
)

type Planet struct {
	box2d.Body
	ID int32
	Rad float64
}

func getMass(d, r float64) float64 {
	area := rl.Pi*(r*r)   // Here area is volume as this is 2d
	return d*area
}

func NewPlanet(id int32, x, y, vx, vy, r, d float64) *Planet {
	p := Planet{
		ID: id,
		Rad: r,
	}
	p.Set(&box2d.Vec2{r*2, r*2}, getMass(d, r))
	p.Body.Position = box2d.Vec2{x, y}
	p.Body.Velocity = box2d.Vec2{vx, vy}
	World.AddBody(&p.Body)
	return &p
}

func (self *Planet) Update() {

}
