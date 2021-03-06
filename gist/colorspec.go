// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gist

import (
	"fmt"
	"image"

	"image/color"

	"github.com/goki/ki/kit"
	"github.com/goki/mat32"
	"github.com/srwiley/rasterx"
)

// Color defines a standard color object for GUI use, with RGBA values, and
// all the usual necessary conversion functions to / from names, strings, etc

// ColorSpec fully specifies the color for rendering -- used in FillStyle and
// StrokeStyle
type ColorSpec struct {
	Source   ColorSources      `desc:"source of color (solid, gradient)"`
	Color    Color             `desc:"color for solid color source"`
	Gradient *rasterx.Gradient `desc:"gradient parameters for gradient color source"`
}

var KiT_ColorSpec = kit.Types.AddType(&ColorSpec{}, nil)

// see colorparse.go for ColorSpec.SetString() method

// ColorSources determine how the color is generated -- used in FillStyle and StrokeStyle
type ColorSources int32

const (
	SolidColor ColorSources = iota
	LinearGradient
	RadialGradient
	ColorSourcesN
)

//go:generate stringer -type=ColorSources

var KiT_ColorSources = kit.Enums.AddEnumAltLower(ColorSourcesN, kit.NotBitFlag, StylePropProps, "")

func (ev ColorSources) MarshalJSON() ([]byte, error)  { return kit.EnumMarshalJSON(ev) }
func (ev *ColorSources) UnmarshalJSON(b []byte) error { return kit.EnumUnmarshalJSON(ev, b) }

// GradientPoints defines points within the gradient
type GradientPoints int32

const (
	GpX1 GradientPoints = iota
	GpY1
	GpX2
	GpY2
	GradientPointsN
)

// IsNil tests for nil solid or gradient colors
func (cs *ColorSpec) IsNil() bool {
	if cs.Source == SolidColor {
		return cs.Color.IsNil()
	}
	return cs.Gradient == nil
}

// ColorOrNil returns the solid color if non-nil, or nil otherwise -- for
// consumers that handle nil colors
func (cs *ColorSpec) ColorOrNil() color.Color {
	if cs.Color.IsNil() {
		return nil
	}
	return cs.Color
}

// SetColor sets a solid color
func (cs *ColorSpec) SetColor(cl color.Color) {
	cs.Color.SetColor(cl)
	cs.Source = SolidColor
	cs.Gradient = nil
}

// SetName sets a solid color by name
func (cs *ColorSpec) SetName(name string) {
	cs.Color.SetName(name)
	cs.Source = SolidColor
	cs.Gradient = nil
}

// Copy copies a gradient, making new copies of the stops instead of
// re-using pointers
func (cs *ColorSpec) CopyFrom(cp *ColorSpec) {
	*cs = *cp
	if cp.Gradient != nil {
		cs.Gradient = &rasterx.Gradient{}
		*cs.Gradient = *cp.Gradient
		sn := len(cp.Gradient.Stops)
		cs.Gradient.Stops = make([]rasterx.GradStop, sn)
		copy(cs.Gradient.Stops, cp.Gradient.Stops)
	}
}

// SetShadowGradient sets a linear gradient starting at given color and going
// down to transparent based on given color and direction spec (defaults to
// "to down")
func (cs *ColorSpec) SetShadowGradient(cl color.Color, dir string) {
	cs.Color.SetColor(cl)
	if dir == "" {
		dir = "to down"
	}
	cs.SetString(fmt.Sprintf("linear-gradient(%v, lighter-0, transparent)", dir), nil)
	cs.Source = LinearGradient
}

// SetGradientBounds sets bounds of the gradient
func SetGradientBounds(grad *rasterx.Gradient, bounds image.Rectangle) {
	grad.Bounds.X = float64(bounds.Min.X)
	grad.Bounds.Y = float64(bounds.Min.Y)
	sz := bounds.Size()
	grad.Bounds.W = float64(sz.X)
	grad.Bounds.H = float64(sz.Y)
}

// CopyGradient copies a gradient, making new copies of the stops instead of
// re-using pointers
func CopyGradient(dst, src *rasterx.Gradient) {
	*dst = *src
	sn := len(src.Stops)
	dst.Stops = make([]rasterx.GradStop, sn)
	copy(dst.Stops, src.Stops)
}

func MatToRasterx(mat *mat32.Mat2) rasterx.Matrix2D {
	return rasterx.Matrix2D{float64(mat.XX), float64(mat.YX), float64(mat.XY), float64(mat.YY), float64(mat.X0), float64(mat.Y0)}
}

// RenderColor gets the color for rendering, applying opacity and bounds for
// gradients
func (cs *ColorSpec) RenderColor(opacity float32, bounds image.Rectangle, xform mat32.Mat2) interface{} {
	if cs.Source == SolidColor || cs.Gradient == nil {
		return rasterx.ApplyOpacity(cs.Color, float64(opacity))
	} else {
		if cs.Source == RadialGradient {
			cs.Gradient.IsRadial = true
		} else {
			cs.Gradient.IsRadial = false
		}
		SetGradientBounds(cs.Gradient, bounds)
		return cs.Gradient.GetColorFunctionUS(float64(opacity), MatToRasterx(&xform))
	}
}

// SetIFace sets the color spec from given interface value, e.g., for ki.Props
// key is an optional property key for error -- always logs errors
func (c *ColorSpec) SetIFace(val interface{}, ctxt Context, key string) error {
	switch valv := val.(type) {
	case string:
		c.SetString(valv, ctxt)
	case *Color:
		c.SetColor(*valv)
	case *ColorSpec:
		*c = *valv
	case color.Color:
		c.SetColor(valv)
	}
	return nil
}
