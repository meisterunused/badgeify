package badgeify

import (
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/template"
)

type badge struct {
	Label       string
	Value       string
	CustomColor string
}

// DefaultColor defines fallback color to lightgrey.
const DefaultColor = "9f9f9f"

// Offset sets a horizontal label and value offset.
const Offset = 10

// ColorRange assigns colors by a certain range limit from 100 to 0 given with:
//
//	above 95 -> "brightgreen"
//	above 90 -> "green"
//	above 85 -> "yellowgreen"
//	above 80 -> "yellow"
//	above 75 -> "orange"
//	above  0 -> "red"
var ColorRange = map[int]string{
	95: "4c1",
	90: "97CA00",
	85: "a4a61d",
	80: "dfb317",
	75: "fe7d37",
	0:  "e05d44",
}

//go:embed badge.svg
var svg []byte

// New creates a new badge by the given label and value.
func New(label, value string) *badge {
	return &badge{label, value, ""}
}

// textLen calculates the width of a given text, we expect one character to be
// the width of 6.
func textLen(text string) int {
	return len(text) * 6
}

// Height returns a static height of the badge.
func (b *badge) Height() int {
	return 20
}

// Width calculates the total width of the badge by its label and value length.
func (b *badge) Width() int {
	return b.LabelLength() + b.ValueLength()
}

// LabelStart provides the center of the label.
func (b *badge) LabelStart() int {
	return b.LabelLength() / 2
}

// LabelLength returns the total
func (b *badge) LabelLength() int {
	return Offset*2 + textLen(b.Label)
}

func (b *badge) ValueStart() int {
	return b.LabelLength() + b.ValueLength()/2
}

func (b *badge) ValueLength() int {
	return Offset*2 + textLen(b.Value)
}

func (b *badge) Color() string {
	if b.CustomColor != "" {
		return b.CustomColor
	}
	color := DefaultColor
	val, err := strconv.Atoi(strings.Trim(b.Value, "%"))
	if err != nil {
		return color
	}
	for limit, limitColor := range ColorRange {
		if val >= limit {
			return limitColor
		}
	}
	return color
}

func (b *badge) ForegroundPath() string {
	return fmt.Sprintf("M0 0h%dv%dH0z", b.Width(), b.Height())
}

func (b *badge) BackgroundPath() string {
	return fmt.Sprintf("M0 0h%dv%dH0z", b.LabelLength(), b.Height())
}

func (b *badge) ColorPath() string {
	mid := b.LabelLength()
	length := b.Width() - mid
	return fmt.Sprintf("M%d 0h%dv%dH%dz", mid, length, b.Height(), mid)
}

func (b *badge) Print(w io.Writer) error {
	funcMap := template.FuncMap{
		"sum": func(nums ...int) int {
			total := 0
			for _, num := range nums {
				total += num
			}
			return total
		},
	}
	tmpl, err := template.New("badge").Funcs(funcMap).Parse(string(svg))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, b)
}
