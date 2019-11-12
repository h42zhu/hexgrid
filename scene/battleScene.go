package scene

import (
	"log"
	"tactics/common"
	"tactics/hexagon"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	imageFiles = []string{
		"./asset/image/enemy.png",
		"./asset/image/spaceship.png",
	}
)

// BattleScene ..
type BattleScene struct {
	Enemy map[string]*hexagon.Entity
	Ally  map[string]*hexagon.Entity
	Grid  *hexagon.HexGrid
	Asset map[string]*pixel.Sprite
}

// Init ..
func (s *BattleScene) Init(r *RenderContext) {
	assetMap, loadFileErr := common.LoadPictures(imageFiles)
	if loadFileErr != nil {
		log.Fatal(loadFileErr)
	}
	s.Asset = assetMap

	// hex grid:
	s.Grid = hexagon.NewHexGrid(40, pixel.V(0, 0), 13, 6)

	// placeholder for testing purposes
	spaceship := hexagon.NewEntity(assetMap["spaceship"], hexagon.NewHexIndex(1, 1), "spaceship", true)
	enemy := hexagon.NewEntity(assetMap["enemy"], hexagon.NewHexIndex(2, 2), "enemy", false)

	s.AddEntity(spaceship)
	s.AddEntity(enemy)
	// s.RenderHexGrid(r)
}

// Render ..
func (s *BattleScene) Render(r *RenderContext) {
	s.RenderAllEntity(r)

}

// UpdateInput ..
func (s *BattleScene) UpdateInput(r *RenderContext) {
	// check for mouse clicks
	if r.win.JustPressed(pixelgl.MouseButtonLeft) {
		// select units
		s.mouseActionClickLeft(r)
	}

	if r.win.JustPressed(pixelgl.MouseButtonRight) {
		// place units
		s.mouseActionClickRight(r)
	}

	s.RenderMousePosition(r, r.win.MousePosition())
}

// NewBattleScene ..
func NewBattleScene() *BattleScene {
	return &BattleScene{
		Enemy: map[string]*hexagon.Entity{},
		Ally:  map[string]*hexagon.Entity{},
	}
}

// AddEntity ..
func (s *BattleScene) AddEntity(entity *hexagon.Entity) {
	if entity.Ally {
		s.Ally[entity.ID] = entity
	} else {
		s.Enemy[entity.ID] = entity
	}
}

// RenderAllEntity ..
func (s *BattleScene) RenderAllEntity(r *RenderContext) {
	for _, entity := range s.Ally {
		RenderEntityHex(r, entity, s.Grid)
	}

	for _, entity := range s.Enemy {
		RenderEntityHex(r, entity, s.Grid)
	}
}

// RenderHexGrid ..
func (s *BattleScene) RenderHexGrid(r *RenderContext) {
	DrawHexGrid(r, s.Grid, 1.0, colornames.Black)
}

// UpdateEntityIdx ..
func (s *BattleScene) UpdateEntityIdx(id string, ally bool, dest hexagon.HexIndex) {
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
func (s *BattleScene) RenderMousePosition(r *RenderContext, pos pixel.Vec) {
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
				DrawHex(r, hexCell, 0, colornames.Aliceblue)
				hg.HoverCell = &idx
				DrawHoverCell(r, hg, 0, colornames.Grey)
			}
		}
	} else {
		hg.HoverCell = &idx
		DrawHoverCell(r, hg, 0, colornames.Grey)
	}
}

func (s *BattleScene) mouseActionClickLeft(r *RenderContext) {
	pos := r.win.MousePosition()
	idx := s.Grid.GetIndex(pos)

	se := s.Grid.SelectedEntity

	if se == "" {
		// no entity selected yet
		for k, v := range s.Ally {
			if v.GetIndex() == idx {
				s.Grid.SelectedEntity = k
				return
			}
		}
	} else if entity, ok := s.Ally[se]; ok {
		// cancel selection
		if entity.GetIndex() != idx {
			s.Grid.SelectedEntity = ""
		}
	} else {
		s.Grid.SelectedEntity = ""
	}
}

func (s *BattleScene) mouseActionClickRight(r *RenderContext) {
	se := s.Grid.SelectedEntity
	if se != "" {
		// move entity to mouse position
		hc := s.Grid.HoverCell
		if entity, ok := s.Ally[se]; ok && hc != nil {
			entity.SetIndex(hexagon.NewHexIndex(hc.X, hc.Y))
		}
	}
}
