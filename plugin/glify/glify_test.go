// +build js

package glify

import (
	"syscall/js"
	"testing"

	"github.com/ctessum/geom"
	"github.com/ctessum/geom/encoding/geojson"
	leaflet "github.com/ctessum/go-leaflet"
)

func TestNewShapes(t *testing.T) {
	doc := js.Global().Get("document")
	elem := doc.Call("createElement", "div")
	elem.Set("width", 100)
	doc.Get("body").Call("appendChild", elem)
	m := leaflet.NewMap(elem, map[string]interface{}{"preferCanvas": true})
	shapes := []geom.Polygon{
		{{{0, 0}, {1, 1}, {0, 1}}},
		{{{1, 0}, {2, 1}, {1, 1}}},
	}
	colors := [][]byte{
		{0, 0, 0},
		{255, 255, 255},
	}

	g := &Geometry{
		Type: "Polygon",
		Features: make([]struct {
			Type     string            `json:"type",js:"type"`
			Geometry *geojson.Geometry `json:"geometry,js:"geometry"`
		}, len(shapes)),
	}
	for i, s := range shapes {
		var err error
		g.Features[i].Geometry, err = geojson.ToGeoJSON(s)
		if err != nil {
			t.Fatal(err)
		}
	}

	colorF := func(i int) (r, g, b float64) {
		bt := colors[i]
		r = float64(uint8(bt[0])) / 255
		g = float64(uint8(bt[1])) / 255
		b = float64(uint8(bt[2])) / 255
		return
	}
	opacity := 1.0
	NewShapes(m, g, colorF, opacity)
}
