package doublebuffer

import (
	"github.com/syndr0m/mygameengine/image"
)

type DoubleBuffer struct {
	width   uint
	height  uint
	buf1    *image.Image
	buf2    *image.Image
	current int8
}

func (d *DoubleBuffer) GetCurrentImage() *image.Image {
	if d.current == 1 {
		return d.buf1
	}
	return d.buf2
}

func (d *DoubleBuffer) GetPreviousImage() *image.Image {
	if d.current == 1 {
		return d.buf2
	}
	return d.buf1
}

func (d *DoubleBuffer) SwapBuffers() {
	d.current = d.current * -1
}

func (d *DoubleBuffer) GetWidth() uint  { return d.width }
func (d *DoubleBuffer) GetHeight() uint { return d.height }

func New(width uint, height uint) *DoubleBuffer {
	doublebuffer := new(DoubleBuffer)
	doublebuffer.width = width
	doublebuffer.height = height
	doublebuffer.current = 1
	doublebuffer.buf1 = image.New(width, height)
	doublebuffer.buf2 = image.New(width, height)
	return doublebuffer
}
