package mygameengine

import (
	"github.com/syndr0m/mygameengine/image"
)

type Board struct {
	frame     uint64
	onStart   func()
	onStop    func()
	onKeyDown func(key int)
	onKeyUp   func(key int)
	onRepaint func(image *image.Image)
}

func (board *Board) OnStart(f func())                     { board.onStart = f }
func (board *Board) OnStop(f func())                      { board.onStop = f }
func (board *Board) OnKeyDown(f func(key int))            { board.onKeyDown = f }
func (board *Board) OnRepaint(f func(image *image.Image)) { board.onRepaint = f }

func (board *Board) GetFrame() uint64 { return board.frame }

func (board *Board) Repaint(image *image.Image) {
	board.frame++
	if board.onRepaint != nil {
		board.onRepaint(image)
	}
}
func (board *Board) Start() {
	board.frame = 0
	if board.onStart != nil {
		board.onStart()
	}
}
func (board *Board) Stop() {
	if board.onStop != nil {
		board.onStop()
	}
}
func (board *Board) KeyDown(key int) {
	if board.onKeyDown != nil {
		board.onKeyDown(key)
	}
}

func NewBoard() *Board {
	board := new(Board)
	return board
}
