package ui

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gosuri/racer/pkg/strutil/color"
)

// Component in the interface that UI components need to implement
type Component interface {
	Format() string
}

// Printer represents the output printer for the ui
type Printer struct {
	// Writer is where the output should writer to
	Writer io.Writer

	// NoColor when true does not display colors
	NoColor bool

	comps []Component
	color *color.Color
}

// NewPrinter returns a pointer to a new printer object
func NewPrinter() *Printer {
	return &Printer{Writer: os.Stdout, color: &color.Color{}}
}

// Add adds the components to the printer
func (p *Printer) Add(c Component) *Printer {
	p.comps = append(p.comps, c)
	return p
}

// AddTitle Adds a Title to the printer
func (p *Printer) AddTitle(title string) *Printer {
	return p.Add(&Title{text: title})
}

// String returns the formmated string of the output
func (p *Printer) String() string {
	var buf bytes.Buffer
	for _, c := range p.comps {
		buf.WriteString(c.Format())
		buf.WriteString("\n")
	}
	return buf.String()
}

// Color returns an instance of color
func (p *Printer) Color() *color.Color {
	if p.color == nil {
		p.color = &color.Color{}
	}
	if p.NoColor {
		p.color.Disable()
	}
	return p.color
}

// Print prints the output to the writer
func (p *Printer) Print() {
	fmt.Fprintln(p.Writer, p.String())
}
