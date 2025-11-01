package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var backgroundColour *ebiten.Image

type Box interface {
	GetBounds() (float64, float64, float64, float64)
}

type Game struct {
	Player       *Player
	frameCounter int
	level        *Level
}

func NewGame() *Game {
	player := NewPlayer(100, 100)

	backgroundColour = ebiten.NewImage(screenWidth, screenHeight)
	backgroundColour.Fill(color.RGBA{128, 128, 255, 255})
	return &Game{
		Player: player,
		level:  NewLevel(),
	}
}

func (g *Game) Update() error {

	g.checkKeys()

	if !g.Player.jumping {
		if math.Abs(g.Player.VelocityX) < 0.001 {
			g.Player.changeState(Idle)
		} else {
			g.Player.changeState(Walking)
		}
	}

	g.Player.Update()

	checkPlatformCollisions(g.Player, g.level.platforms)

	if g.Player.Y > 500 {
		g.Player.Y = 500
		g.Player.VelocityY = 0
	}

	g.frameCounter++
	if (g.frameCounter)%g.Player.AnimationSpeed == 0 {
		g.Player.IncFrame()
		g.frameCounter = 0
	}

	return nil
}

func (g *Game) checkKeys() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.VelocityX = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Player.VelocityX = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && !g.Player.jumping {
		g.Player.VelocityY = -5
		g.Player.jumping = true
		g.Player.changeState(JumpingUp)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.level.draw(screen)
	g.Player.Draw(screen)

	debugString := fmt.Sprintf("%s, x: %.2f, y: %.2f, w: %.2f, h: %.2f", g.Player.playerState.String(), g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height)
	ebitenutil.DebugPrint(screen, debugString)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkPlatformCollisions(player *Player, platforms []Platform) {
	player.OnGround = false
	eX, eY, eW, eH := player.GetBounds()
	for _, platform := range platforms {
		pX, pY, pW, pH := platform.GetBounds()

		if player.VelocityY > 0 {
			if eX < pX+pW &&
				eX+eW > pX &&
				eY < pY+pH &&
				eY+eH > pY {

				player.Y = pY - eH
				player.jumping = false
				player.OnGround = true
				player.VelocityY = 0
				player.Friction = platform.friction
			}
		}
	}
}
