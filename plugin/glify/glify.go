// +build js

// Package glify is a wrapper for the glify javascipt LeafletJS plugin,
// for use with gopherjs.
package glify

import (
	"sync"
	"syscall/js"

	"github.com/norunners/vert"

	"github.com/ctessum/geom/encoding/geojson"
	leaflet "github.com/ctessum/go-leaflet"
)

func initialize() {
	leaflet.Initialize()

	doc := js.Global().Get("document")

	glifyJSBytes, err := Asset("js/glify.js")
	if err != nil {
		panic(err)
	}

	// Load glify javascript.
	script := doc.Call("createElement", "script")
	script.Set("type", "text/javascript")
	script.Set("text", string(glifyJSBytes))
	doc.Get("head").Call("appendChild", script)
}

var initializeOnce sync.Once

// Initialize this package dy loading the leaflet JS and CSS.
func Initialize() {
	initializeOnce.Do(initialize)
}

// Shapes is a wrapper for the shapes type.
type Shapes struct {
	js.Value
}

// NewShapes returns a new Shapes variable.
// colors should return the color of the shape at index i, where each color channel is in the range [0,1].
func NewShapes(m *leaflet.Map, shapes *Geometry, colors func(i int) (r, g, b float64), opacity float64) *Shapes {
	Initialize()
	jsShapes := vert.ValueOf(shapes)
	colorFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		r, g, b := colors(args[0].Int())
		return map[string]interface{}{"r": r, "g": g, "b": b}
	})
	options := map[string]interface{}{
		"map":     m.Value,
		"data":    jsShapes,
		"color":   colorFunc,
		"opacity": opacity,
	}
	return &Shapes{
		Value: leaflet.L.Get("glify").Call("shapes", options),
	}
}

type Geometry struct {
	Type     string `json:"type",js:"type",`
	Features []struct {
		Type     string            `json:"type",js:"type"`
		Geometry *geojson.Geometry `json:"geometry,js:"geometry"`
	} `json:"features",js:"features"`
}
