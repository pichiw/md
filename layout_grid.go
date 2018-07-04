package md

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

// LayoutGrid creates a new layout grid
func LayoutGrid(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		append(
			[]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-layout-grid"),
				),
			},
			markup...,
		)...,
	)
}

// LayoutGridInner creates a new layout grid inner element
func LayoutGridInner(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		append(
			[]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-layout-grid__inner"),
				),
			},
			markup...,
		)...,
	)
}

type LayoutGridCellOptions struct {
	Span int
}

// LayoutGridCell creates a new layout grid inner element
func LayoutGridCell(opts LayoutGridCellOptions, markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		append(
			[]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-layout-grid__cell", Span("mdc-layout-grid__cell", opts.Span)),
				),
			},
			markup...,
		)...,
	)
}
