package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	backgroundImage *ebiten.Image
	platforms       []Platform
	gravity         float64
}

// func NewLevel(backgroundImage *ebiten.Image, platforms []Platform, gravity float64) *Level {
func NewLevel() *Level {

	l := &Level{}

	l.backgroundImage = ebiten.NewImage(screenWidth, screenHeight)
	l.backgroundImage.Fill(color.RGBA{128, 128, 255, 255})

	basePlatform := NewPlatform(0, 500, float64(screenWidth), 5, 0.99999, color.RGBA{0, 255, 0, 255})

	smallPlatform := NewPlatform(300, 450, 50, 5, 0.1, color.RGBA{255, 0, 0, 255})

	l.platforms = []Platform{basePlatform, smallPlatform}
	l.gravity = 0.2
	return l
}

func (l *Level) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(0, 0)
	screen.DrawImage(l.backgroundImage, op)

	for _, platform := range l.platforms {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(platform.X, platform.Y)
		screen.DrawImage(platform.image, op)
	}

}
