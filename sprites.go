package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprites struct {
	WalkSprites     []*ebiten.Image
	IdleSprites     []*ebiten.Image
	JumpUpSprites   []*ebiten.Image
	JumpDownSprites []*ebiten.Image
}

func splitSpriteSheet(img image.Image, row, cells int) []*ebiten.Image {
	var sprites []*ebiten.Image
	cellWidth := 32
	cellHeight := 32

	for i := 0; i < cells; i++ {
		x0 := i * cellWidth
		y0 := row * cellHeight
		x1 := x0 + cellWidth
		y1 := y0 + cellHeight
		subImg := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(x0, y0, x1, y1))
		pixeledImage := actualImage(subImg)
		sprites = append(sprites, ebiten.NewImageFromImage(pixeledImage))
	}
	return sprites
}

func (p *Player) IncFrame() {
	p.CurrentFrame++
	if p.CurrentFrame >= len(p.CurrentSprites) {
		p.CurrentFrame = 0
	}
}

// walk row 3 (8 images)
// jump row 5 (8 images)
// idle row 0 (2 images)
func (p *Sprites) loadImages() {
	// Load player images for walking, idle, and jumping
	file, err := os.Open("assets/AnimationSheet_Character.png") // todo: make it work embedded
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.WalkSprites = splitSpriteSheet(img, 3, 8)
	p.IdleSprites = splitSpriteSheet(img, 0, 2)
	p.IdleSprites = append(p.IdleSprites, splitSpriteSheet(img, 1, 2)...)
	jumpSprites := splitSpriteSheet(img, 5, 8)

	// block := ebiten.NewImage(32, 32)
	// block.Fill(image.Black)

	// p.IdleSprites = []*ebiten.Image{block}

	p.JumpUpSprites = jumpSprites[:4]
	p.JumpDownSprites = jumpSprites[4:]

}

func actualImage(img image.Image) image.Image {
	bounds := img.Bounds()
	minX, minY := bounds.Max.X, bounds.Max.Y
	maxX, maxY := bounds.Min.X, bounds.Min.Y

	nothingFound := false

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y).(color.NRGBA)
			if c.A > 0 { // Check if the pixel is not fully transparent
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
				nothingFound = true
			}
		}
	}

	if !nothingFound {
		return img
	}
	return img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(minX-1, minY, maxX+1, maxY))
}
