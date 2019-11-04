package main

import (
	"fmt"
	"log"
	"tactics/common"
	"tactics/control"
	"tactics/hexagon"

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
		Bounds: pixel.R(0, 0, 800, 600),
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
	hg := hexagon.NewHexGrid(40, pixel.V(0, 0), 4)

	renderer := hexagon.NewRenderer(win, imd, basicAtlas)
	renderer.DrawHexGrid(&hg, 1, colornames.Green)

	assetMap, loadFileErr := common.LoadPictures(imageFiles)
	if loadFileErr != nil {
		log.Fatal(err)
	}
	fmt.Println(assetMap)

	// mouse control
	mc := &control.MouseControl{
		Win: win,
	}

	win.SetSmooth(true)
	// last := time.Now()

	// main game loop
	for !win.Closed() {
		// time tick
		// dt := time.Since(last).Seconds()
		// last = time.Now()

		// check mouse input
		if win.JustPressed(pixelgl.MouseButtonLeft) {

			mc.SelectCell(&hg)
			renderer.DrawHexGrid(&hg, 1, colornames.Green)
			renderer.DrawSelectedCell(&hg, 1, colornames.Hotpink)
		}

		// update
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		renderer.DrawTextGrid(&hg, 1, colornames.Black)

		win.Update()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
