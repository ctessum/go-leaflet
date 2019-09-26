// +build js

// Package leaflet provides a (currently minimal) wrapper around leaflet.js
// for use with gopherjs. The bindings are currently for leaflet version 1.5.1.
package leaflet

import (
	"sync"
	"syscall/js"
)

// L is the primary leaflet javascript object.
var L js.Value

func initialize() {
	doc := js.Global().Get("document")

	// Load leaflet CSS.
	link := doc.Call("createElement", "link")
	link.Set("href", "https://unpkg.com/leaflet@1.5.1/dist/leaflet.css")
	link.Set("type", "text/css")
	link.Set("rel", "stylesheet")
	doc.Get("head").Call("appendChild", link)

	// Load leaflet javascript.
	script := doc.Call("createElement", "script")
	script.Set("src", "https://unpkg.com/leaflet@1.5.1/dist/leaflet.js")
	doc.Get("head").Call("appendChild", script)

	var wg sync.WaitGroup
	wg.Add(1)
	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		L = js.Global().Get("L")
		callback.Release()
		wg.Done()
		return nil
	})
	script.Set("onreadystatechange", callback)
	script.Set("onload", callback)
	wg.Wait()
}

var initializeOnce sync.Once

// Initialize this package dy loading the leaflet JS and CSS.
func Initialize() {
	initializeOnce.Do(initialize)
}

// Map is a leaflet map object: http://leafletjs.com/reference-1.5.1.html#map
type Map struct {
	js.Value
}

// NewMap creates a new map in the specified div with the specified options.
// Possible options are listed at: https://leafletjs.com/reference-1.5.0.html#map-factory.
func NewMap(div js.Value, options map[string]interface{}) *Map {
	if div == js.Null() {
		panic("leaflet: cannot use null map div")
	}
	if nodeName := div.Get("nodeName").String(); nodeName != "DIV" {
		panic("leaflet: map div nodeName should be DIV but is " + nodeName)
	}
	Initialize()
	return &Map{
		Value: L.Call("map", div, options),
	}
}

// SetView sets the center and zoom level of the map.
func (m *Map) SetView(center *LatLng, zoom int) {
	m.Value.Call("setView", center.Value, zoom)
}

// CreatePane creates a new Pane with the given name:
// http://leafletjs.com/reference-1.5.0.html#map-createpane
func (m *Map) CreatePane(name string) *Pane {
	return &Pane{Value: m.Value.Call("createPane", name)}
}

// Pane is a leaflet pane.
type Pane struct {
	js.Value
}

// SetZIndex sets the Z index of the pane.
func (p *Pane) SetZIndex(index int) {
	p.Value.Get("style").Set("zIndex", index)
}

// LatLng specifies a point in latitude and longitude
type LatLng struct {
	js.Value
}

// NewLatLng returns a new LatLng object.
func NewLatLng(lat, lng float64) *LatLng {
	Initialize()
	return &LatLng{
		Value: L.Call("latLng", lat, lng),
	}
}

// TileLayer is a leaflet TileLayer object: http://leafletjs.com/reference-1.5.0.html#tilelayer
type TileLayer struct {
	Layer
}

// NewTileLayer creates a new TileLayer with the specified URL template and
// options.
func NewTileLayer(urlTemplate string, options map[string]interface{}) *TileLayer {
	Initialize()
	return &TileLayer{
		Layer: Layer{
			Value: L.Call("tileLayer", urlTemplate, options),
		},
	}
}

// Layer is a leaflet layer object: http://leafletjs.com/reference-1.5.0.html#layer.
type Layer struct {
	js.Value
}

// AddTo add the receiver to the specified Map.
func (l *Layer) AddTo(m *Map) {
	l.Value.Call("addTo", m)
}

// Path is a leaflet path object: http://leafletjs.com/reference-1.5.0.html#path.
type Path struct {
	Layer
}

// SetStyle sets the style of the receiver:
// http://leafletjs.com/reference-1.5.0.html#path-setstyle.
func (p *Path) SetStyle(style map[string]interface{}) {
	p.Value.Call("setStyle", style)
}

// Polyline is a leaflet polyline object: http://leafletjs.com/reference-1.5.0.html#polyline.
type Polyline struct {
	Path
}

// Polygon is a leaflet polygon object: http://leafletjs.com/reference-1.5.0.html#polygon.
type Polygon struct {
	Polyline
}

// NewPolygon creates a new polygon.
func NewPolygon(latlngs []*LatLng) *Polygon {
	Initialize()
	return &Polygon{
		Polyline: Polyline{
			Path: Path{
				Layer: Layer{
					Value: L.Call("polygon", latlngs),
				},
			},
		},
	}
}

// GridLayer is a leaflet GridLayer: http://leafletjs.com/reference-1.5.0.html#gridlayer.
type GridLayer struct {
	Layer
}

// NewGridLayer creates a new GridLayer.
func NewGridLayer() *GridLayer {
	Initialize()
	return &GridLayer{
		Layer: Layer{
			Value: L.Call("gridLayer"),
		},
	}
}
