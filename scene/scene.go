package scene

import (
	"tactics/hexagon"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

// Scene ..
type Scene struct {
	Enemy    map[pixel.Vec]*Entity
	Ally     map[pixel.Vec]*Entity
	Renderer *Renderer
	Grid     *hexagon.HexGrid
}

// Entity ..
type Entity struct {
	Sprite *pixel.Sprite
	Index  pixel.Vec
	Type   string
	Ally   bool
}

// NewEntity returns a new entity
func NewEntity(s *pixel.Sprite, idx pixel.Vec, t string, ally bool) *Entity {
	return &Entity{
		Sprite: s,
		Index:  idx,
		Type:   t,
		Ally:   ally,
	}
}

// NewScene ..
func NewScene(r *Renderer, hg *hexagon.HexGrid) *Scene {
	return &Scene{
		Enemy:    map[pixel.Vec]*Entity{},
		Ally:     map[pixel.Vec]*Entity{},
		Renderer: r,
		Grid:     hg,
	}
}

// AddEntity ..
func (s *Scene) AddEntity(entity *Entity) {
	if entity.Ally {
		s.Ally[entity.Index] = entity
	} else {
		s.Enemy[entity.Index] = entity
	}
}

// RenderAllEntity ..
func (s *Scene) RenderAllEntity() {
	for _, entity := range s.Ally {
		s.Renderer.RenderEntityHex(entity, s.Grid)
	}

	for _, entity := range s.Enemy {
		s.Renderer.RenderEntityHex(entity, s.Grid)
	}
}

// RenderHexGrid ..
func (s *Scene) RenderHexGrid() {
	s.Renderer.DrawHexGrid(s.Grid, 1.0, colornames.Black)
}

// RenderMousePosition ..
func (s *Scene) RenderMousePosition(pos pixel.Vec) {
	hg := s.Grid
	idx := hg.GetIndex(pos)

	if _, ok := hg.Cells[idx]; !ok {
		return
	}

	if hg.HoverCell != nil {
		prevHoverCell := *(hg.HoverCell)
		// if val, ok := dict["foo"]; ok
		if prevHoverCell != idx {
			if _, ok := hg.Cells[prevHoverCell]; ok {
				hexCell := hg.Cells[prevHoverCell]
				s.Renderer.DrawHex(hexCell, 0, colornames.Aliceblue)
				hg.HoverCell = &idx
				s.Renderer.DrawHoverCell(hg, 0, colornames.Grey)
			}
		}
	} else {
		hg.HoverCell = &idx
		s.Renderer.DrawHoverCell(hg, 0, colornames.Grey)
	}
}
