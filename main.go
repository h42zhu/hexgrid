package main

import (
	"fmt"
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

	// grid system
	// grid := util.NewGridRender(5, 5, 50, pixel.Vec{X: 0, Y: 0}, win, imd, basicAtlas)
	// grid.ShowGrid()

	hg := hexagon.NewHexGrid(40, pixel.V(0, 0), 4)

	renderer := hexagon.NewRenderer(win, imd, basicAtlas)
	renderer.DrawHexGrid(hg, 1, colornames.Green)

	// mouse control
	// mc := &util.MouseControl{
	// 	Win:        win,
	// 	GridRender: grid,
	// }

	// win.SetSmooth(true)
	// last := time.Now()

	// main game loop
	for !win.Closed() {
		// time tick
		// dt := time.Since(last).Seconds()
		// last = time.Now()

		// check mouse input
		// mc.Update(dt)

		// update grid

		win.Clear(colornames.Aliceblue)
		imd.Draw(win)

		win.Update()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
