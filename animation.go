package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const ANIMATION_DEFAULT_SPEED = 5

type Animation struct {
	Speed int
	Frames []*ebiten.Image
}

func AnimationNew() *Animation {
	a := Animation{}

	a.Speed = ANIMATION_DEFAULT_SPEED

	return &a
}

func (a *Animation) GetFrame(g *Game) *ebiten.Image {
	if len(a.Frames) == 0 {
		return ebiten.NewImage(0, 0)
	}

	i := (g.Count / a.Speed) % len(a.Frames)

	return a.Frames[i]
}

func (a *Animation) PushFrame(f *ebiten.Image) {
	a.Frames = append(a.Frames, f)
}