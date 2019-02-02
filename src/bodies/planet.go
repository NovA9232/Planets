package bodies

import (
	phy "github.com/gen2brain/raylib-go/physics"
	"github.com/gen2brain/raylib-go/raylib"
)

type Planet struct {
	id int
	r float32  // Radius
	d float32  // Densitya
	currForce rl.Vector2
	body *phy.Body
}

func NewPlanet(id int, x, y int, r, d float32) {
	p := new(Planet)
	p.Init(id, x, y, r, d)
	ListOfBods = append(ListOfBods, p)
}

func (p *Planet) Init(id int, x, y int, r, d float32) {
	p.id = id
	p.body = phy.NewBodyCircle(rl.NewVector2(float32(x), float32(y)), r, d)
	p.r = r
	p.d = d
}

func (p *Planet) SetForce(fx, fy float32) {
	dx, dy := fx-p.currForce.X, fy-p.currForce.Y
	p.body.AddForce(rl.NewVector2(dx, dy))
}

func (p *Planet) getBody() *phy.Body {
	return p.body
}
