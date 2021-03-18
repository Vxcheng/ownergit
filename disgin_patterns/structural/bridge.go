package structural

import "fmt"

type Drawing interface {
	DoDraw(name string) string
}

type circleBridge struct {
	Color
	Frame
	category string
}

func NewCircleDrawing(c Color, f Frame) Drawing {
	return &circleBridge{
		Frame:    f,
		Color:    c,
		category: "circle",
	}
}

func (c *circleBridge) DoDraw(name string) string {
	return fmt.Sprintf("%s now drawing %s use %s color , it is %s frame", name, c.category, c.color(), c.metirail())
}

// color shape
type Color interface {
	color() string
}

type redColor struct{}

func NewRedColor() Color {
	return &redColor{}
}

func (c *redColor) color() string {
	return "red"
}

type greenColor struct{}

func NewGreenColor() Color {
	return &greenColor{}
}

func (c *greenColor) color() string {
	return "green"
}

type Frame interface {
	metirail() string
}

type wooden struct{}

func NewWooden() Frame {
	return &wooden{}
}

func (w *wooden) metirail() string {
	return "wooden"
}
