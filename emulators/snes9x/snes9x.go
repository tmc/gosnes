package snes9x

import (
	"context"

	"github.com/tmc/gosnes/emulators"
)

// Emulator is the SNES9X emulator.
type Emulator struct{}

var _ emulators.Emulator = &Emulator{}

// NewEmulator prepares a new Emulator.
func NewEmulator(ctx context.Context) (*Emulator, error) {
	return &Emulator{}, nil
}
