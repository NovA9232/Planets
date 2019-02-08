package planet

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (self *Planet) debugVel() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Position.X), float32(self.Position.Y)), rl.NewVector2(float32(self.Position.X+self.Velocity.X), float32(self.Position.Y+self.Velocity.Y)), 2, rl.Green)
}

func (self *Planet) debugForce() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Position.X), float32(self.Position.Y)), rl.NewVector2(float32(self.Position.X+self.Force.X/1000), float32(self.Position.Y+self.Force.Y/1000)), 2, rl.Red)
}
