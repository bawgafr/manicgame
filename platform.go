package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	Entity
	friction float64
	image    *ebiten.Image
	barrier  bool
	direction
}

func newBarrier(x, y, width, height, friction float64, colour color.Color, d direction) Platform {
	p := NewPlatform(x, y, width, height, friction, colour, d)
	p.barrier = true
	return p
}

func NewPlatform(x, y, width, height, friction float64, colour color.Color, d direction) Platform {

	// need to scale the width and height from the level coords to the screen coords

	var w, h float64
	switch d {
	case horiz:
		w = 800
		h = 5.0
	case vert:
		w = 5.0
		h = 450
	case literal:
		w = width
		h = height
	}

	platformImg := ebiten.NewImage(int(w), int(h))
	platformImg.Fill(colour)

	return Platform{
		Entity: Entity{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		friction:  friction,
		image:     platformImg,
		direction: d,
	}
}
