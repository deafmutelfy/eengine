package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Camera struct {
	Viewport f64.Vec2
	Position f64.Vec2

	Scale float64
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.Viewport[0] * 0.5,
		c.Viewport[1] * 0.5,
	}
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}