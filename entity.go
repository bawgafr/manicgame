package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X              float64
	Y              float64
	VelocityX      float64
	VelocityY      float64
	Width          float64
	Height         float64
	CurrentSprites []*ebiten.Image
	CurrentFrame   int
	AnimationSpeed int
}

func (e Entity) GetBounds() (float64, float64, float64, float64) {
	return e.X, e.Y, e.Width, e.Height
}
