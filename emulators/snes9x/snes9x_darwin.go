package snes9x

import "fmt"

// Launch launches or finds reference to a running emulator.
func (s *Emulator) Launch() error {
	fmt.Println("Launching")
	return nil
}
