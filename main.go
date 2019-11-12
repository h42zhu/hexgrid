package main

import (
	"fmt"
	"log"
	"tactics/scene"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 960, 720),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// draw shapes
	imd := imdraw.New(nil)
	// text altas
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// render context
	renderCtx := scene.NewRenderContext(win, imd, basicAtlas)

	// battle scene
	battleScene := scene.NewBattleScene()

	// scene manager
	sceneManager := scene.NewManager(battleScene, renderCtx)

	win.SetSmooth(true)

	// main game loop
	for !win.Closed() {
		win.Clear(colornames.Aliceblue)

		sceneManager.UpdateInput()
		sceneManager.Render()
	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
