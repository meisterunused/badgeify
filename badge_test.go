package badgeify

import (
	_ "embed"
	"testing"
)

func TestLabelStart(t *testing.T) {
	res := New("label", "value").LabelStart()
	if res != 25 {
		t.Errorf("Got label-start at %d, want 25", res)
	}
}

func TestValueStart(t *testing.T) {
	res := New("label", "value").ValueStart()
	if res != 75 {
		t.Errorf("Got value-start at %d, want 75", res)
	}
}

func TestLabelLength(t *testing.T) {
	ln := New("test", "").LabelLength()
	if ln != 44 {
		t.Errorf("LabelLength(test) = %d, want 44", ln)
		return
	}
	ln = New("long test", "").LabelLength()
	if ln != 74 {
		t.Errorf("LabelLength(long test) = %d, want 74", ln)
	}
}

func TestValueLength(t *testing.T) {
	ln := New("", "80%").ValueLength()
	if ln != 38 {
		t.Errorf("ValueLength(test) = %d, want 38", ln)
		return
	}
	ln = New("", "100%").ValueLength()
	if ln != 44 {
		t.Errorf("ValueLength(long test) = %d, want 44", ln)
	}
}

func TestDefaultColor(t *testing.T) {
	color := New("works", "fine").Color()
	if color != DefaultColor {
		t.Errorf("Got default color #%s, want #%s", color, DefaultColor)
	}
}

func TestCustomColor(t *testing.T) {
	badge := New("works", "fine")
	badge.CustomColor = "ababab"
	if badge.Color() != "ababab" {
		t.Errorf("Got custom color #%s, want #ababab", badge.Color())
	}
}

func TestColorRange(t *testing.T) {
	color := New("ok", "89").Color()
	if color != ColorRange[85] {
		t.Errorf("Got color #%s, want > 85", color)
	}
	color = New("not good", "75").Color()
	if color != ColorRange[75] {
		t.Errorf("Got color #%s, want > 75", color)
	}
	color = New("bad", "60").Color()
	if color != ColorRange[0] {
		t.Errorf("Got color #%s, want > 0", color)
	}
}

func TestForegroundPath(t *testing.T) {
	path := New("coverage", "60%").ForegroundPath()
	if path != "M0 0h106v20H0z" {
		t.Errorf("Got SVG foreground path %s, want M0 0h106v20H0z", path)
	}
}

func TestBackgroundPath(t *testing.T) {
	path := New("coverage", "60%").BackgroundPath()
	if path != "M0 0h68v20H0z" {
		t.Errorf("Got SVG background path %s, want M0 0h68v20H0z", path)
	}
}

func TestColorPath(t *testing.T) {
	path := New("coverage", "60%").ColorPath()
	if path != "M68 0h38v20H68z" {
		t.Errorf("Got SVG background path %s, want M68 0h38v20H68z", path)
	}
}
