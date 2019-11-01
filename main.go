package main

import (
	"fmt"
	"tactics/control"
	"tactics/hexagon"

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
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// draw shapes
	imd := imdraw.New(nil)
	// text altas
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// hex grid system
	hg := hexagon.NewHexGrid(40, pixel.V(0, 0), 4)

	renderer := hexagon.NewRenderer(win, imd, basicAtlas)
	renderer.DrawHexGrid(&hg, 1, colornames.Green)

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

			mouse := mc.GetMousePosition()
			fmt.Println(mouse)
			idx := hg.GetIndex(mouse)
			fmt.Println(idx)

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
