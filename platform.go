package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	Entity
	friction float64
	image    *ebiten.Image
}

func NewPlatform(x, y, width, height, friction float64, colour color.Color) Platform {

	platformImg := ebiten.NewImage(int(width), int(height))
	platformImg.Fill(colour)

	return Platform{
		Entity: Entity{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		friction: friction,
		image:    platformImg,
	}
}
