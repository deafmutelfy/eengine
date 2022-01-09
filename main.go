package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	ebiten.SetWindowSize(SW, SH)
	ebiten.SetWindowTitle("Cock and ball shooter")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)

	check(ebiten.RunGame(GameNew()), true)
}