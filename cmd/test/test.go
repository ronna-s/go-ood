package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

//go:embed gopher.jpg
var gopher []byte

func drewImage(s tcell.Screen, img image.Image, x1, y1, x2, y2 int) {
	img, err := jpeg.Decode(bytes.NewReader(gopher))
	if err != nil {
		return
	}
	rect := img.Bounds()
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			//r, g, b, _ := c.RGBA()
			//r2, g1, b = uint8(r>>8), uint8(g>>8), uint8(b>>8)
			//n := new(big.Int)
			//hex := fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
			//n.SetString(hex, 16)
			color := tcell.FromImageColor(img.At(x, y))
			style := tcell.StyleDefault.Background(color)
			s.SetContent(x, y, tcell.RuneBoard, nil, style)
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	img, err := jpeg.Decode(bytes.NewReader(gopher))
	if err != nil {
		return
	}
	//log.Println(c)
	//r, g, b, _ := c.RGBA()
	//style = tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(b)))

	rect := img.Bounds()
	xRatio := rect.Dx() / (x2 - x1)
	yRatio := rect.Dy() / (y2 - y1)
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			style = tcell.StyleDefault.Background(tcell.FromImageColor(img.At(x*xRatio, y*yRatio)))
			s.SetContent(x, y, ' ', nil, style)
		}
	}
	//
	////style = tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(b)))
	//if y2 < y1 {
	//	y1, y2 = y2, y1
	//}
	//if x2 < x1 {
	//	x1, x2 = x2, x1
	//}
	//
	//// Fill background
	//for row := y1; row <= y2; row++ {
	//	for col := x1; col <= x2; col++ {
	//		s.SetContent(col, row, ' ', nil, style)
	//	}
	//}
	//
	//// Draw borders
	//for col := x1; col <= x2; col++ {
	//	s.SetContent(col, y1, tcell.RuneHLine, nil, style)
	//	s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	//}
	//for row := y1 + 1; row < y2; row++ {
	//	s.SetContent(x1, row, tcell.RuneVLine, nil, style)
	//	s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	//}
	//
	//// Only draw corners if necessary
	//if y1 != y2 && x1 != x2 {
	//	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	//	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	//	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	//	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	//}
	//
	//drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func main() {
	//defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	//s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	drawBox(s, 1, 1, 60, 30, boxStyle, "Click and drag to draw a box")
	//drawBox(s, 5, 9, 32, 300, boxStyle, "Press C to reset")

	// Event loop
	ox, oy := -1, -1
	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			button := ev.Buttons()
			// Only process button events, not wheel events
			button &= tcell.ButtonMask(0xff)

			if button != tcell.ButtonNone && ox < 0 {
				ox, oy = x, y
			}
			switch ev.Buttons() {
			case tcell.ButtonNone:
				if ox >= 0 {
					label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
					drawBox(s, ox, oy, x, y, boxStyle, label)
					ox, oy = -1, -1
				}
			}
		}
	}
}
