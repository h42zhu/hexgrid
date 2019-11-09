package main

import (
	"fmt"
	"log"
	"tactics/common"
	"tactics/control"
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
	battleScene := scene.NewScene(renderer, hg)
	// battleScene.RenderHexGrid()

	assetMap, loadFileErr := common.LoadPictures(imageFiles)
	if loadFileErr != nil {
		log.Fatal(err)
	}

	spaceship := hexagon.NewEntity(assetMap["spaceship"], hexagon.NewHexIndex(1, 1), "spaceship", true)
	enemy := hexagon.NewEntity(assetMap["enemy"], hexagon.NewHexIndex(2, 2), "enemy", false)

	battleScene.AddEntity(spaceship)
	battleScene.AddEntity(enemy)

	// mouse control
	mc := control.NewMouseControl(win)

	win.SetSmooth(true)

	// main game loop
	for !win.Closed() {
		// clear
		win.Clear(colornames.Aliceblue)

		// check mouse input
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			// select units
			mc.MouseActionClickLeft(battleScene)
		}

		if win.JustPressed(pixelgl.MouseButtonRight) {
			// place units
			mc.MouseActionClickRight(battleScene)
		}

		battleScene.RenderMousePosition(win.MousePosition())

		// battleScene.RenderHexGrid()

		imd.Draw(win)
		battleScene.RenderAllEntity()

		win.Update()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
