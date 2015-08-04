package mygameengine

import (
	"fmt"
	"github.com/syndr0m/mygameengine/doublebuffer"
	"github.com/syndr0m/mygameengine/image"
	"time"
)

type MyGameEngine struct {
	screen *doublebuffer.DoubleBuffer
	fps    uint
	frame  uint64
	boards *Boards
	assets *Assets
	exit   chan int
}

type LoopFunc func(*MyGameEngine)

func (engine *MyGameEngine) Run() {
	// using gxui to open the window
	commands := make(chan Message)
	events := make(chan Message)
	go gxuiOpenWindow(engine.GetScreenImage().GetWidth(), engine.GetScreenImage().GetHeight(), engine.GetScreen(), commands, events)

	// engine-loop rendering
	ticker := time.NewTicker(time.Millisecond * time.Duration(1000/engine.fps))
	go func() {
		for {
			<-ticker.C
			// repaint
			engine.frame++
			// game repaint code
			engine.Boards().Current().Repaint(engine.GetScreenImage())
			// switching buffer
			engine.GetScreen().SwapBuffers()
			// saving screen on the double buffer
			commands <- Message{MESSAGE_REPAINT, 0}
		}
	}()
	// engine-loop keys
	go func() {
		for {
			var m Message = <-events
			fmt.Println("EVENT MESSAGE RECEIVED")
			switch m.name {
			case MESSAGE_KEY_DOWN:
				fmt.Println("EVENT MESSAGE DISPATCHED TO KEYDOWN")
				engine.Boards().Current().KeyDown(m.value)
			case MESSAGE_EXIT:
				engine.Stop()
			}
		}
	}()

	engine.exit = make(chan int)
	<-engine.exit
}

func (engine *MyGameEngine) Stop() {
	engine.exit <- 42
}

func (engine *MyGameEngine) GetFrame() uint64 {
	return engine.frame
}

func (engine *MyGameEngine) GetFps() uint {
	return engine.fps
}

func (engine *MyGameEngine) GetScreen() *doublebuffer.DoubleBuffer {
	return engine.screen
}

func (engine *MyGameEngine) GetScreenImage() *image.Image {
	return engine.screen.GetCurrentImage()
}

func (engine *MyGameEngine) Boards() *Boards { return engine.boards }
func (engine *MyGameEngine) Assets() *Assets { return engine.assets }

//
func New(screenWidth uint, screenHeight uint, fps uint) *MyGameEngine {
	engine := new(MyGameEngine)
	engine.boards = NewBoards()
	engine.assets = NewAssets()
	// default screen.
	engine.screen = doublebuffer.New(screenWidth, screenHeight)
	engine.fps = fps
	return engine
}
