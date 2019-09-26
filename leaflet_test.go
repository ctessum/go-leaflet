// +build js

package leaflet

import (
	"syscall/js"
	"testing"
)

func TestNewMap(t *testing.T) {
	doc := js.Global().Get("document")
	elem := doc.Call("createElement", "div")
	NewMap(elem, map[string]interface{}{"preferCanvas": true})
}
