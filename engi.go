// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engi

import (
	"log"
	"math"
	"time"
)

var (
	responder Responder
	config    *Config
	timing    *stats
	Files     *Loader
	GL        *gl2
)

type Action int
type Key int
type Modifier int

var (
	MOVE    = Action(0)
	PRESS   = Action(1)
	RELEASE = Action(2)
	SHIFT   = Modifier(0x0001)
	CONTROL = Modifier(0x0002)
	ALT     = Modifier(0x0004)
	SUPER   = Modifier(0x0008)
)

// A Config holds settings for your game's window and application.
type Config struct {
	// Title is the name of the created window.
	// Default: Untitled
	Title string

	// Width and Height are hints about the size of the window. You
	// may not end up with the indicated size, so you should always
	// query eng for the true width and height after initialization.
	// Default: 1024 x 640
	Width  int
	Height int

	// Fullscreen tells eng whether to open windowed or fullscreen.
	// Default: false
	Fullscreen bool

	// Vsync enables or disables vertical sync which will limit the
	// number of frames rendered per second to your monitor's refresh
	// rate. This may or may not be supported on certain platforms.
	// Default: true
	Vsync bool

	// Resizable tells eng if it should request a window that can be
	// resized by the user of your game.
	// Default: false
	Resizable bool

	// Fsaa indicates how many samples to use for the multisampling
	// buffer. Generally it will be 1, 2, 4, 8, or 16.
	// Default: 1
	Fsaa int

	// PrintFPS turns on a logging of the frames per second to the
	// console every second.
	// Default: false
	LogFPS bool
}

func NewConfig() *Config {
	return &Config{"ENG!", 800, 600, false, true, false, 1, false}
}

type Responder interface {
	init()
	draw()
	resize(width, height int)
	Preload()
	Setup()
	Update()
	Mouse(x, y float32, action Action)
	Scroll(amount float32)
	Key(key Key, modifier Modifier, action Action)
	Type(char rune)
}

type stats struct {
	Elapsed, Dt, Fps, Frames, Period float64
	Start                            time.Time
	Then                             time.Time
	show                             bool
}

func NewStats(show bool) *stats {
	st := new(stats)
	st.Start = time.Now()
	st.Period = 1
	st.Update()
	st.show = show
	return st
}

func (t *stats) Update() {
	now := time.Now()
	t.Frames += 1
	t.Dt = now.Sub(t.Then).Seconds()
	t.Elapsed += t.Dt
	t.Then = now

	if t.Elapsed >= t.Period {
		t.Fps = t.Frames / t.Period
		t.Elapsed = math.Mod(t.Elapsed, t.Period)
		t.Frames = 0
		if t.show {
			log.Println(t.Fps)
		}
	}
}

// Run should be called with a type that satisfies the Responder
// interface. Windows will be setup using your Config and a runloop
// will start, blocking the main thread and calling methods on the
// given responder.
func Open(title string, width, height int, fullscreen bool, r Responder) {
	OpenConfig(&Config{title, width, height, fullscreen, true, false, 1, false}, r)
}

func OpenConfig(c *Config, r Responder) {
	config = c
	responder = r
	Files = NewLoader()
	run()
}