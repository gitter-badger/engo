package main

import (
	"image/color"
	"log"

	"github.com/paked/engi"
	"github.com/paked/engi/ecs"
)

type Game struct{}

func (game *Game) Preload() {
	engi.Files.Add("assets/326488.wav")
}

func (game *Game) Setup(w *ecs.World) {
	engi.SetBg(color.White)

	w.AddSystem(&engi.RenderSystem{})
	w.AddSystem(&engi.AudioSystem{})

	backgroundMusic := ecs.NewEntity([]string{"AudioSystem"})
	backgroundMusic.AddComponent(&engi.AudioComponent{File: "326488.wav", Repeat: true, Background: true})

	err := w.AddEntity(backgroundMusic)
	if err != nil {
		log.Println(err)
	}
}

func (*Game) Hide()        {}
func (*Game) Show()        {}
func (*Game) Type() string { return "Game" }

func main() {
	opts := engi.RunOptions{
		Title:  "Audio Demo",
		Width:  1024,
		Height: 640,
	}
	engi.Run(opts, &Game{})
}
