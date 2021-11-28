// Command extractsprites pulls out sprites from an snes9x image (with backgrounds hidden).
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"os"

	"github.com/tmc/gosnes/utilities/iterm2helpers"
)

var (
	flagShowBoundingBox = flag.Bool("show-bb", false, "If true, print out the bounding box that most closely matches the input dataset.")
)

func main() {
	flag.Parse()
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(files []string) error {
	fmt.Println("running for", len(files), "images")

	// 1. compose all images
	// 2. find best-fit bounding box
	// 3. crop each unique input to bounding box and remove background

	combined, err := composeImages(files)
	if err != nil {
		return fmt.Errorf("issue composing images: %w", err)
	}

	if err := iterm2helpers.PrintGIFToTerminal(combined); err != nil {
		return fmt.Errorf("issue printing image: %w", err)
	}
	o, err := os.OpenFile("output.gif", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	gif.EncodeAll(o, combined)
	o.Close()

	return nil
}

func composeImages(files []string) (*gif.GIF, error) {
	result := &gif.GIF{}

	// frames := []image.Image{}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		img, err := png.Decode(f)
		if err != nil {
			return nil, err
		}
		const defaultDelay = 10
		pal := color.Palette(palette.Plan9)
		pbounds := image.Rect(0, 0, 46, 64)
		ebounds := image.Rect(240, 180, 250, 320)
		pimg := image.NewPaletted(pbounds, pal)
		for y := 0; y < pbounds.Max.Y; y++ {
			for x := 0; x < pbounds.Max.X; x++ {
				pimg.Set(x, y, img.At(x+ebounds.Min.X, y+ebounds.Min.Y))
			}
		}
		// pimg = pimg.SubImage(bounds).(*image.Paletted)
		result.Image = append(result.Image, pimg)
		result.Delay = append(result.Delay, defaultDelay)
	}

	return result, nil
}
