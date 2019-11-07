package main

import (
	"fmt"
	"log"
	"tactics/common"
	"tactics/hexagon"
	"tactics/scene"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	imageFiles = []string{
		"./asset/image/enemy.png",
		"./asset/image/spaceship.png",
	}
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

	// hex grid system
	hg := hexagon.NewHexGrid(40, pixel.V(0, 0), 13, 6)

	renderer := scene.NewRenderer(win, imd, basicAtlas)
	battleScene := scene.NewScene(renderer)
	battleScene.RenderHexGrid(hg)

	assetMap, loadFileErr := common.LoadPictures(imageFiles)
	if loadFileErr != nil {
		log.Fatal(err)
	}

	spaceship := scene.NewEntity(assetMap["spaceship"], pixel.V(1, 1), "spaceship", true)
	enemy := scene.NewEntity(assetMap["enemy"], pixel.V(2, 2), "enemy", false)

	battleScene.AddEntity(spaceship)
	battleScene.AddEntity(enemy)

	// mouse control
	// mc := &control.MouseControl{
	// 	Win: win,
	// }

	win.SetSmooth(true)

	// main game loop
	for !win.Closed() {
		// clear
		win.Clear(colornames.Aliceblue)

		// check mouse input
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			idx := hg.GetIndex(win.MousePosition())
			fmt.Println(idx)
		}
		battleScene.RenderMousePosition(hg, win.MousePosition())

		battleScene.RenderHexGrid(hg)

		imd.Draw(win)
		battleScene.RenderAllEntity(hg)

		win.Update()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
