package mygameengine

import (
	"image/color"
)

var COLOR_WHITE color.RGBA = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var COLOR_BLACK color.RGBA = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var COLOR_RED color.RGBA = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var COLOR_GREEN color.RGBA = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var COLOR_BLUE color.RGBA = color.RGBA{R: 0, G: 0, B: 255, A: 255}

func CreateColor(r uint8, g uint8, b uint8, a uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: a}
}
