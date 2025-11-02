package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var debugMode = false

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

	level := NewLevel(screenWidth, screenHeight)
	player := NewPlayer(100, 100, level)

	backgroundColour = ebiten.NewImage(screenWidth, screenHeight)
	backgroundColour.Fill(color.RGBA{128, 128, 255, 255})
	return &Game{
		Player: player,
		level:  level,
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

	if !g.Player.FeetOnGround {
		g.Player.VelocityY += 0.2
		if g.Player.VelocityY > 5 {
			g.Player.VelocityY = 5
		}
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
		g.Player.VelocityY = g.Player.JumpForce
		g.Player.jumping = true
		g.Player.changeState(JumpingUp)
		g.Player.FeetOnGround = false
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	// should do the background image here....?

	g.level.draw(screen)
	g.Player.Draw(screen, *g.level)

	// debugString :=
	ebitenutil.DebugPrint(screen, g.Player.String())

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkPlatformCollisions(player *Player, platforms []Platform) {

	eX, eY, eW, eH := player.GetFeetBounds()

	onHorizPlatform := false

	for _, platform := range platforms {
		switch platform.direction {
		case horiz, literal:
			pX, pY, pW, pH := platform.GetBounds()

			if player.VelocityY >= 0 {

				entityLeft := eX
				entityRight := eX + eW
				platformLeft := pX
				platformRight := pX + pW

				if (entityLeft > platformLeft && entityLeft < platformRight) ||
					(entityRight > platformLeft && entityRight < platformRight) {

					entityTop := eY + eH
					entityBottom := eY

					platformTop := pY - 2
					platformBottom := pY + pH + 2

					if (entityTop > platformTop && entityTop < platformBottom) ||
						(entityBottom > platformTop && entityBottom < platformBottom) {
						// then the feet are within the box of the platform

						// amend the player's y (player.Y) so that they are standing on the top
						// of the platform (which should also register as on plafrom)
						player.Y = platformTop - (player.Scale.Y * player.Height)

						player.jumping = false
						player.VelocityY = 0
						player.Friction = platform.friction
						onHorizPlatform = true
						player.FeetOnGround = true
					}
				}
			}
		case vert:
			// sort out the walls later
		}

	}
	if !onHorizPlatform {
		player.FeetOnGround = false
	}
}
