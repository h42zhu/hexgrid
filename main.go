package main

import (
	"fmt"
	"tactics/hexagon"
	"time"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"

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
	// basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// grid system
	// grid := util.NewGridRender(5, 5, 50, pixel.Vec{X: 0, Y: 0}, win, imd, basicAtlas)
	// grid.ShowGrid()

	// hex test
	hg := hexagon.NewHexGrid(20, pixel.V(400, 300), 3)

	renderer := hexagon.NewRenderer(win, imd)
	renderer.DrawHexGrid(hg, 0, colornames.Green)

	// mouse control
	// mc := &util.MouseControl{
	// 	Win:        win,
	// 	GridRender: grid,
	// }

	win.SetSmooth(true)
	last := time.Now()

	// main game loop
	for !win.Closed() {
		// time tick
		dt := time.Since(last).Seconds()
		last = time.Now()

		fmt.Println(dt)

		// check mouse input
		// mc.Update(dt)

		// update grid
		// grid.ShowGrid(colornames.Green, colornames.Red)

		win.Clear(colornames.Aliceblue)
		imd.Draw(win)

		win.Update()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
