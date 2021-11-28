package iterm2helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
)

// PrintImageToTerminal prints an image in iTerm2 format.
func PrintImageToTerminal(img image.Image) error {
	var buf = new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return err
	}
	printOSC()
	fmt.Printf("1337;File=inline=1;")
	fmt.Printf(":%s", base64.StdEncoding.EncodeToString(buf.Bytes()))
	printST()
	fmt.Println()
	return nil
}

// PrintGIFToTerminal prints a GIF in iTerm2 format.
func PrintGIFToTerminal(g *gif.GIF) error {
	var buf = new(bytes.Buffer)
	if err := gif.EncodeAll(buf, g); err != nil {
		return err
	}
	printOSC()
	fmt.Printf("1337;File=inline=1;")
	fmt.Printf(":%s", base64.StdEncoding.EncodeToString(buf.Bytes()))
	printST()
	fmt.Println()
	return nil
}

// BEL is the bell character.
const BEL = "\a"

// ESC is the escape character.
const ESC = "\033"

func printOSC() {
	if os.Getenv("TERM") == "screen" {
		fmt.Printf(ESC + `Ptmux;` + ESC + ESC + `]`)
	} else {
		fmt.Printf(ESC + "]")
	}
}

func printST() {
	if os.Getenv("TERM") == "screen" {
		fmt.Printf(BEL + ESC + `\`)
	} else {
		fmt.Printf(BEL)
	}
}
