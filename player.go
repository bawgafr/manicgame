package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type playerState int

const (
	Idle playerState = iota
	Walking
	JumpingUp
	JumpingDown
)

func (ps playerState) String() string {
	switch ps {
	case Idle:
		return "Idle"
	case Walking:
		return "Walking"
	case JumpingUp:
		return "JumpingUp"
	case JumpingDown:
		return "JumpingDown"
	default:
		return "Unknown"
	}
}

type Player struct {
	Entity
	Sprites
	playerState  playerState // controls sprites etc
	Friction     float64     // set when on a platform
	jumping      bool        // jump key has been pressed.... or only allow jump when on ground...
	FeetOnGround bool
	Scale        Scale
	screenScale  Scale
	JumpForce    float64
}

func (p Player) String() string { // debug string
	px, py, pw, ph := p.GetBounds()
	fx, fy, fw, fh := p.GetFeetBounds()
	return fmt.Sprintf("%s, x: %.2f, y: %.2f, w: %.2f, h: %.2f\nfx: %.2f, fy: %.2f, fw: %.2f, fh: %.2f", p.playerState.String(), px, py, pw, ph, fx, fy, fw, fh)
}

type Scale struct {
	X float64
	Y float64
}

func NewPlayer(x, y float64, level *Level) *Player {

	p := &Player{
		Entity: Entity{
			X:              x,
			Y:              y,
			CurrentFrame:   0,
			AnimationSpeed: 20,
		},
		Scale: Scale{X: 2, Y: 2},

		playerState:  JumpingDown,
		jumping:      false, // jump key has been pressed.... or only allow jump when on ground...
		FeetOnGround: false,
		JumpForce:    -8.0,
	}

	ssx := coordMapResize(p.Scale.X, 0, level.Width, 0, screenWidth)
	ssy := coordMapResize(p.Scale.Y, 0, level.Height, 0, 450)
	p.screenScale = Scale{X: ssx, Y: ssy}

	fmt.Println("Player screen scale set to:", p.screenScale.X, p.screenScale.Y)

	// get the sprites from the asset file for the player
	p.loadImages()

	// sprites are all much thinner than the sprite sheet cell size, so we need to adjust
	// width and height accordingly
	allSprites := append(append(p.WalkSprites, p.IdleSprites...), append(p.JumpUpSprites, p.JumpDownSprites...)...)
	w, h := maxWidthHeight(allSprites)

	p.Entity.Width = float64(w)
	p.Entity.Height = float64(h)

	p.changeState(p.playerState)
	return p
}

func (p *Player) Update() {
	// apply gravity

	//oldXVel := p.VelocityX
	//oldYVel := p.VelocityY

	// update the position based on velocity
	p.X += p.VelocityX
	p.Y += p.VelocityY

	if p.FeetOnGround {
		// apply friction only when on ground
		p.VelocityX *= p.Friction
	} else {
		// air resistance
		p.VelocityX *= 0.99
	}

	if math.Abs(p.VelocityX) < 0.4 {
		p.changeState(Idle)
	}

	if !p.FeetOnGround && p.VelocityY > 0 {
		p.changeState(JumpingDown)
	}

}

func (p *Player) changeState(state playerState) {
	if p.playerState != state {
		p.playerState = state
		p.CurrentFrame = 0
		switch state {
		case Idle:
			p.CurrentSprites = p.IdleSprites
			p.Entity.AnimationSpeed = 20
		case Walking:
			p.CurrentSprites = p.WalkSprites
			p.Entity.AnimationSpeed = 10
		case JumpingUp:
			p.CurrentSprites = p.JumpUpSprites
			p.Entity.AnimationSpeed = 30

		case JumpingDown:
			p.CurrentSprites = p.JumpDownSprites
			p.Entity.AnimationSpeed = 30
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image, level Level) {

	px, py, _, _ := p.GetBounds()
	x, y := XYCoordMap(px, py, &level)
	if debugMode {
		fx, fy, fw, fh := p.GetFeetBounds()
		fx, fy = XYCoordMap(fx, fy, &level)

		vector.FillRect(screen, float32(fx), float32(fy), float32(fw), float32(fh), color.RGBA{255, 255, 255, 10}, true)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	if p.VelocityX < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(p.Width, 0)
	}
	op.GeoM.Scale(p.screenScale.X, p.screenScale.Y)
	op.GeoM.Translate(x, y)

	screen.DrawImage(p.CurrentSprites[p.CurrentFrame], op)

}

func (p Player) GetBounds() (float64, float64, float64, float64) {
	return p.X, p.Y, p.Scale.X * p.Width, p.Scale.Y * p.Height
}

func (p Player) GetFeetBounds() (float64, float64, float64, float64) {
	return p.X, p.Y + (p.Scale.Y * p.Height) + 5, p.Scale.X * p.Width, -10.0
}
