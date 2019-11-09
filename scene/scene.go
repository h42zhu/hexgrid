package scene

import (
	"tactics/hexagon"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

// Scene ..
type Scene struct {
	Enemy    map[string]*hexagon.Entity
	Ally     map[string]*hexagon.Entity
	Renderer *Renderer
	Grid     *hexagon.HexGrid
}

// NewScene ..
func NewScene(r *Renderer, hg *hexagon.HexGrid) *Scene {
	return &Scene{
		Enemy:    map[string]*hexagon.Entity{},
		Ally:     map[string]*hexagon.Entity{},
		Renderer: r,
		Grid:     hg,
	}
}

// AddEntity ..
func (s *Scene) AddEntity(entity *hexagon.Entity) {
	if entity.Ally {
		s.Ally[entity.ID] = entity
	} else {
		s.Enemy[entity.ID] = entity
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

// UpdateEntityIdx ..
func (s *Scene) UpdateEntityIdx(id string, ally bool, dest hexagon.HexIndex) {
	if ally {
		if entity, ok := s.Ally[id]; ok {
			entity.SetIndex(dest)
		}
	} else {
		if entity, ok := s.Enemy[id]; ok {
			entity.SetIndex(dest)
		}
	}
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
