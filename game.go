package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	collisions "github.com/tducasse/ebiten-collisions"
	"golang.org/x/image/math/f64"
)

const SW = 640
const SH = 480

type Game struct{
	Count int
	Sprites map[string]*Sprite

	World *collisions.World
	Camera *Camera
	buff *ebiten.Image
}

func (g *Game) GetSpriteNameById(id string) string {
	for k, v := range g.Sprites {
		if v.Id == id {
			return k
		}
	}

	return ""
}

func GameNew() *Game {
	g := Game{}
	
	g.Sprites = make(map[string]*Sprite)
	g.World = collisions.MakeWorld()
	g.Camera = &Camera{Viewport: f64.Vec2{SW, SH}}
	g.buff = ebiten.NewImage(SW * 3, SH * 3)

	g.Sprites["eggplant"] = SpriteNew()
	g.Sprites["eggplant"].Animations["default"] = AnimationNew()
	g.Sprites["eggplant"].CurrentAnimation = "default"
	frame, _,  err := ebitenutil.NewImageFromFile("assets/eggplant.png")
	check(err, true)
	g.Sprites["eggplant"].Animations["default"].PushFrame(frame)
	g.Sprites["eggplant"].X = 60
	g.Sprites["eggplant"].Y = 60
	g.Sprites["eggplant"].Scale = 6
	g.Sprites["eggplant"].SetupCollisions(g.World, 12, 12)

	g.Sprites["fedora"] = SpriteNew()
	g.Sprites["fedora"].Animations["default"] = AnimationNew()
	g.Sprites["fedora"].CurrentAnimation = "default"
	fframe, _,  err := ebitenutil.NewImageFromFile("assets/fedora.png")
	check(err, true)
	g.Sprites["fedora"].Animations["default"].PushFrame(fframe)
	g.Sprites["fedora"].X = 100
	g.Sprites["fedora"].Y = 100
	g.Sprites["fedora"].Scale = 0.5
	g.Sprites["fedora"].SetupCollisions(g.World, 205, 205)
	g.Sprites["fedora"].CollisionCallback = func(b1, b2 *collisions.Box) bool {
		t := b2.Data.(string)

		log.Printf("found collision with %s\n", g.GetSpriteNameById(t))

		return true
	}

	log.Println(g.Sprites)

	return &g
}

func (g *Game) Update() error {
	g.Count++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.buff.Clear()
	g.buff.Fill(color.White)

	sp := 1.0

	if ebiten.IsKeyPressed(ebiten.KeyRight)  {
		g.Sprites["fedora"].X += sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Sprites["fedora"].X -= sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Sprites["fedora"].Y -= sp
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Sprites["fedora"].Y += sp
	}

	g.Sprites["eggplant"].Draw(g.buff, g)
	g.Sprites["fedora"].Draw(g.buff, g)

	g.Camera.Position[0] = (g.Sprites["fedora"].X - SW / 2) + (float64(g.Sprites["fedora"].W) * g.Sprites["fedora"].Scale / 2)
	g.Camera.Position[1] = (g.Sprites["fedora"].Y - SH / 2) + (float64(g.Sprites["fedora"].W) * g.Sprites["fedora"].Scale / 2)

	g.Camera.Render(g.buff, screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d\n", int(ebiten.CurrentFPS())))
}

func (g *Game) Layout(_, _ int) (int, int) {
	return SW, SH
}