package main

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = mustLoadImage("assets/Players/bunny1_ready.png")
var Background = loadBackground("assets/Background/bg_layer1.png")

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
}

//go:embed assets/*
var assets embed.FS

func loadBackground(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bg := &ebiten.DrawImageOptions{}
	screen.DrawImage(Background, bg)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	//op.GeoM.Scale(0, 0)
	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition: Vector{X: 100, Y: 100},
	}

	ebiten.SetWindowSize(1080, 1080)
	ebiten.SetWindowTitle("Your game's title")
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
