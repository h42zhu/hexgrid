package common

import (
	"github.com/dyrkin/fsm"
)

// Unit represent a combat unit in game
type Unit struct {
	Morale float64
	Player bool
	State  fsm.State
}
