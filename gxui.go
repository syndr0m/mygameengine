package mygameengine

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/syndr0m/mygameengine/doublebuffer"
)

func gxuiOpenWindow(width uint, height uint, dblBuf *doublebuffer.DoubleBuffer, commands chan Message, events chan Message) {
	gl.StartDriver(func(driver gxui.Driver) {
		theme := flags.CreateTheme(driver)
		window := theme.CreateWindow(int(width), int(height), "MyGameEngine")
		window.SetScale(flags.DefaultScaleFactor)
		screen := theme.CreateImage()
		window.AddChild(screen)
		window.OnClose(func() {
			driver.Terminate()
			events <- Message{MESSAGE_EXIT, 0}
		})
		window.OnKeyDown(func(e gxui.KeyboardEvent) {
			fmt.Println("keydown") // FIXME: without this line, randomly crash ...
			events <- Message{MESSAGE_KEY_DOWN, int(e.Key)}
		})

		// repaint function
		go func() {
			for {
				<-commands
				last := screen.Texture()
				driver.CallSync(func() {
					texture := driver.CreateTexture(dblBuf.GetPreviousImage().GetBuffer(), 1)
					screen.SetTexture(texture)
					if last != nil {
						last.Release()
					}
				})
			}
		}()
	})
}
