package main

import (
	"fmt"

	"github.com/faiface/pixel"
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
	win.SetSmooth(true)
	// last := time.Now()

	// main game loop
	for !win.Closed() {
		// dt := time.Since(last).Seconds()

		// last = time.Now()

	}
}

func main() {
	fmt.Println("start...")
	pixelgl.Run(run)
}
