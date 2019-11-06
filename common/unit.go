package common

import (
	"github.com/dyrkin/fsm"
)

// UnitInfo represent a combat unitinfo in game
type UnitInfo struct {
	Morale float64
	Ally   bool
	State  fsm.State
}
