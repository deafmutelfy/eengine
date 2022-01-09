package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/segmentio/ksuid"
	collisions "github.com/tducasse/ebiten-collisions"
)

type Sprite struct {
	Id string

	X float64
	Y float64

	W int
	H int

	Scale float64

	Animations map[string]*Animation
	CurrentAnimation string

	CollisionCallback func(*collisions.Box, *collisions.Box) bool

	world *collisions.World
	box *collisions.Box
}

func SpriteNew() *Sprite {
	s := Sprite{}
	
	s.Id = ksuid.New().String()
	s.Scale = 1.0
	s.Animations = make(map[string]*Animation)

	return &s
}

func (s *Sprite) Draw(screen *ebiten.Image, g *Game) {
	if s.CurrentAnimation == "" {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X / s.Scale, s.Y / s.Scale)
	op.GeoM.Scale(s.Scale, s.Scale)

	s.box.X = s.X
	s.box.Y = s.Y

	s.world.Move(s.box, 0, 0, s.CollisionCallback)
	
	f := s.Animations[s.CurrentAnimation].GetFrame(g)
	s.W, s.H = f.Size()

	screen.DrawImage(f, op)
}

func (s *Sprite) SetupCollisions(world *collisions.World, W int, H int) {
	s.world = world

	s.box = collisions.MakeBox(s.X / s.Scale, s.Y / s.Scale, float64(W) * s.Scale, float64(H) * s.Scale)
	s.box.AddData(s.Id)

	s.world.Add(s.box)
}