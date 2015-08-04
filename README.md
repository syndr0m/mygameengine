Description
===========

small game engine in golang 
this engine is a small poc to test/learn using golang

Features
========

- image manipulation (works on Windows 8.x & Windows 10) using gxui
- double buffering
- assets (png) loading

Usage
=====

```go
package main

import (
	"fmt"
	"github.com/syndr0m/mygameengine"
)

func main() {
	var screenWidth uint = 640
	var screenHeight uint = 480
	var fps uint = 25

	engine := mygameengine.New(screenWidth, screenHeight, fps)
	intro := mygameengine.NewBoard()
	intro.OnKeyDown(func(key int) { fmt.Println("KEY DOWN", key) })
	intro.OnRepaint(func(screen *image.Image) {
		screen.DrawRectangle(0, 0, screenWidth, screenHeight, mygameengine.COLOR_RED)
	})
	engine.Boards().SetCurrent(intro)
	engine.Run()
}
```

Install
=======

Prerequisites
-------------

you need go, git and mercurial installed

Installing dependencies
-----------------------

Installing the gxui

```
go get -u github.com/google/gxui/...
```

Installing mygameengine

```
go get -u github.com/syndr0m/mygameengine/...
```

Build
-----

```
go build
```
 
# Games

[gotris](https://github.com/syndr0m/gotris)
