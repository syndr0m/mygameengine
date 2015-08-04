package mygameengine

import (
	"github.com/syndr0m/mygameengine/image"
)

type Assets struct {
	pngs map[string]*image.Image
}

func (assets *Assets) Png(name string) {
	assets.pngs[name], _ = image.Png(name)
}

func (assets *Assets) Get(name string) *image.Image {
	return assets.pngs[name]
}

func NewAssets() *Assets {
	assets := new(Assets)
	assets.pngs = make(map[string]*image.Image)
	return assets
}
