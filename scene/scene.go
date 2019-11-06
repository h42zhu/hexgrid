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
func NewScene(r *Renderer) *Scene {
	return &Scene{
		Enemy:    map[pixel.Vec]*Entity{},
		Ally:     map[pixel.Vec]*Entity{},
		Renderer: r,
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
func (s *Scene) RenderAllEntity(hg *hexagon.HexGrid) {
	for _, entity := range s.Ally {
		s.Renderer.RenderEntityOnHexGrid(entity, hg)
	}

	for _, entity := range s.Enemy {
		s.Renderer.RenderEntityOnHexGrid(entity, hg)
	}
}

// RenderHexGrid ..
func (s *Scene) RenderHexGrid(hg *hexagon.HexGrid) {
	s.Renderer.DrawHexGrid(hg, 1.0, colornames.Black)
}

// RenderMousePosition ..
func (s *Scene) RenderMousePosition(hg *hexagon.HexGrid, pos pixel.Vec) {
	idx := hg.GetIndex(pos)
	if _, ok := hg.Cells[idx]; !ok {
		return
	}

	if hg.SelectedCell != nil {
		prevSelectedCell := *(hg.SelectedCell)
		// if val, ok := dict["foo"]; ok
		if prevSelectedCell != idx {
			if _, ok := hg.Cells[prevSelectedCell]; ok {
				hexCell := hg.Cells[prevSelectedCell]
				s.Renderer.DrawHex(hexCell, 0, colornames.White)
				hg.SelectedCell = &idx
				s.Renderer.DrawSelectedCell(hg, 0, colornames.Grey)
			}
		}
	} else {
		hg.SelectedCell = &idx
		s.Renderer.DrawSelectedCell(hg, 0, colornames.Grey)
	}
}
