package md

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

// Button creates a new button
func Button(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Button(
		append([]vecty.MarkupOrChild{
			vecty.Markup(vecty.Attribute("class", "foo-button mdc-button")),
		}, markup...)...,
	)
}
