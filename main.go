package main

import "github.com/hajimehoshi/ebiten/v2"

// a manic minor style game.
// player moving left and right with jumps
// a level with platforms to begin with
// when that works, get moving platforms.

// player effected by gravity. not falling through platforms.

// levels can be larger than the screen. camera follows player.

// uses ebiten library for graphics and input handling.

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {


	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Manic Minor Style Game")
	// ebiten.SetTPS(60)
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
