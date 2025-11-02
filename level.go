package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	backgroundImage *ebiten.Image
	platforms       []Platform
	gravity         float64
	Width           float64
	Height          float64
}

type direction int

const (
	vert direction = iota
	horiz
	literal
)

// func NewLevel(backgroundImage *ebiten.Image, platforms []Platform, gravity float64) *Level {
func NewLevel(w, h float64) *Level {

	l := basicLevel(w, h)

	l.backgroundImage = ebiten.NewImage(screenWidth, screenHeight)
	l.backgroundImage.Fill(color.RGBA{128, 128, 255, 255})

	smallPlatform := NewPlatform(300, 450, 50, 5, 0.1, color.RGBA{255, 0, 0, 255}, literal)

	l.platforms = append(l.platforms, smallPlatform)
	l.gravity = 0.2
	return l
}

func basicLevel(width, height float64) *Level {

	topBarrier := newBarrier(0, 0, width, 5, 0.8, color.RGBA{0, 255, 0, 255}, horiz)
	bottomBarrier := newBarrier(0, height-5, width, 5, 0.8, color.RGBA{0, 255, 0, 255}, horiz)

	leftBarrier := newBarrier(0, 0, 5, height, 0.8, color.RGBA{0, 255, 0, 255}, vert)
	rightBarrier := newBarrier(width-5, 0, 5, height, 0.8, color.RGBA{0, 255, 0, 255}, vert)

	l := &Level{
		Width:     width,
		Height:    height,
		platforms: []Platform{bottomBarrier, topBarrier, leftBarrier, rightBarrier},
	}

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
		x, y := XYCoordMap(platform.X, platform.Y, l)

		op.GeoM.Translate(x, y)
		screen.DrawImage(platform.image, op)
	}

}
