// Package glify is a wrapper for the glify javascipt LeafletJS plugin,
// for use with gopherjs.
package glify

import (
	leaflet "github.com/ctessum/go-leaflet"
	"github.com/gopherjs/gopherjs/js"
)

// Shapes is a wrapper for the shapes type.
type Shapes struct {
	Object *js.Object
}

// NewShapes returns a new Shapes variable.
func NewShapes(options *ShapeOptions) *Shapes {
	return &Shapes{
		Object: leaflet.L.Get("glify").Call("shapes", options),
	}
}

// ShapeOptions holds the options for creating a new Shapes object
// and should be initialized with DefaultShapeOptions.
type ShapeOptions struct {
	Object  *js.Object
	Map     *leaflet.Map         `js:"map"`
	Data    *js.Object           `js:"data"`
	Color   func(index int) *RGB `js:"color"`
	Opacity float64              `js:"opacity"` // number [0,1]
}

// DefaultShapeOptions returns the default ShapeOptions.
func DefaultShapeOptions() *ShapeOptions {
	return &ShapeOptions{
		Object: js.Global.Get("Object").New(),
	}
}

// RGB hold color information and should be initialized with NewRGB.
type RGB struct {
	Object *js.Object
	R      float64 `js:"r"`
	G      float64 `js:"g"`
	B      float64 `js:"b"`
}

// NewRGB returns a new RGB variable.
func NewRGB() *RGB {
	return &RGB{
		Object: js.Global.Get("Object").New(),
	}
}
