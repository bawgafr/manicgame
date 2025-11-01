package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
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
	playerState playerState
	Friction    float64
	jumping     bool
	OnGround    bool
	Sprites
	Entity
	ScaleX float64
	ScaleY float64
}

func NewPlayer(x, y float64) *Player {

	p := &Player{
		Entity: Entity{
			X:              x,
			Y:              y,
			Width:          32,
			Height:         32,
			CurrentFrame:   0,
			AnimationSpeed: 20,
		},
		ScaleX:      2,
		ScaleY:      2,
		playerState: Idle,
		jumping:     false,
		OnGround:    true,
	}
	p.loadImages()

	allSprites := append(append(p.WalkSprites, p.IdleSprites...), append(p.JumpUpSprites, p.JumpDownSprites...)...)

	// set width and height to max of all sprites
	w, h := maxWidthHeight(allSprites)
	p.Entity.Width = float64(w)
	p.Entity.Height = float64(h)
	p.changeState(p.playerState)
	p.Entity.CurrentSprites = p.IdleSprites
	return p
}

func maxWidthHeight(sprites []*ebiten.Image) (int, int) {
	maxW, maxH := 0, 0
	for _, img := range sprites {
		w := img.Bounds().Dx()
		h := img.Bounds().Dy()
		if w > maxW {
			maxW = w
		}
		if h > maxH {
			maxH = h
		}
	}
	return maxW, maxH
}

func (p *Player) Update() {
	// apply gravity
	p.VelocityY += 0.2

	// oldXVel := p.VelocityX
	oldYVel := p.VelocityY

	// update the position based on velocity
	p.X += p.VelocityX
	p.Y += p.VelocityY

	if p.OnGround {
		// apply friction only when on ground
		p.VelocityX *= p.Friction
	} else {
		// air resistance
		p.VelocityX *= 0.99
	}
	p.VelocityX *= 0.9

	// when jump, vy goes to -2
	// gravity adds +ve amount to it.
	// hits top and velocity is 0
	// then velocity goes +ve
	// till we hit ground again.
	// when we hit ground, set jumping to false.

	// how to tell though...

	if oldYVel < 0 && p.VelocityY >= 0 {
		p.changeState(JumpingDown)
	}

	if p.Y == 150 {
		p.jumping = false
	}

	// if math.Abs(p.VelocityY) < 0.0001 { // fails when we hit the top and velocity is essentially zero.
	// 	p.jumping = false
	// }

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

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	if p.VelocityX < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(p.Width, 0)
	}
	op.GeoM.Scale(p.ScaleX, p.ScaleY)
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(p.CurrentSprites[p.CurrentFrame], op)
}

func (p Player) GetBounds() (float64, float64, float64, float64) {
	return p.X, p.Y, 2 * p.Width, 2 * p.Height
}
