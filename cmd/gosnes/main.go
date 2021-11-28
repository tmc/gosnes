package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/tmc/gosnes/emulators"
	"github.com/tmc/gosnes/emulators/snes9x"
)

var (
	flagDriver      = flag.String("emulator", "snes9x", "The emulator to run and manage")
	flagGameProfile = flag.String("gameprofile", "alinktothepast", "The emulator to run and manage")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := run(ctx, *flagDriver, *flagGameProfile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, emulatorName string, gameProfileName string) error {
	var (
		e   emulators.Emulator
		err error
	)
	switch emulatorName {
	case "snes9x":
		e, err = snes9x.NewEmulator(ctx)
	default:
		return fmt.Errorf("unsupported emulator %s", emulatorName)
	}
	if err != nil {
		return fmt.Errorf("issue initializing emulator '%s': %w", emulatorName, err)
	}

	gameProfile, ok := GameProfiles[gameProfileName]

	if !ok {
		return fmt.Errorf("unkown game profile '%s", gameProfileName)
	}

	return gameProfile(e)
}

// GameProfiles contains rom-specific extraction logic.
var GameProfiles = map[string]func(emulators.Emulator) error{
	"alinktothepast": func(e emulators.Emulator) error {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			return err
		}
		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}

		if err := cmdTab(); err != nil {
			return err
		}

		// Or you can use Press and Release
		press := func(key int) {
			kb.SetKeys(key)
			kb.Press()
			fmt.Println("pressed")
			time.Sleep(100 * time.Millisecond)
			screenshot()
			time.Sleep(100 * time.Millisecond)
			screenshot()
			time.Sleep(100 * time.Millisecond)
			screenshot()
			time.Sleep(100 * time.Millisecond)
			screenshot()
			time.Sleep(100 * time.Millisecond)
			screenshot()
			kb.Release()
			fmt.Println("released")
			time.Sleep(10 * time.Millisecond)
		}
		circle := func() {
			for i := 0; i < 3; i++ {
				press(keybd_event.VK_LEFT)
			}
			for i := 0; i < 3; i++ {
				press(keybd_event.VK_UP)
			}
			for i := 0; i < 3; i++ {
				press(keybd_event.VK_RIGHT)
			}
			for i := 0; i < 3; i++ {
				press(keybd_event.VK_DOWN)
			}
		}

		for i := 0; i < 3; i++ {
			circle()
		}

		return nil
	},
}

func cmdTab() error {
	cmd := exec.Command("cmdtab.sh")
	return cmd.Run()
}

func screenshot() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	kb.SetKeys(keybd_event.VK_GRAVE)
	return kb.Launching()
}
