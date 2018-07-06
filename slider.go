package md

import (
	"fmt"

	"github.com/gowasm/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

type Slider struct {
	vecty.Core

	Value int
}

func (s *Slider) Mount() {
	doc := js.Global().Get("document")

	root := doc.Call("getElementById", "slider-example")
	sliderEl := root.Call("querySelector", "#continuous-mdc-slider")
	slider := js.Global().Get("mdc").Get("slider").Get("MDCSlider").New(sliderEl)
	foundation := slider.Get("foundation_")
	slider.Call("listen", "MDCSlider:change",
		js.NewCallback(
			func(vs []js.Value) {
				fmt.Println("New Value:", foundation.Get("value_"))
			},
		),
	)
}

func (s *Slider) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
			vecty.Markup(
				prop.ID("slider-example"),
			),
			elem.Div(
				vecty.Markup(
					prop.ID("continuous-mdc-slider"),
					vecty.Class("mdc-slider"),
					vecty.Attribute("tabindex", "0"),
					vecty.Attribute("role", "slider"),
					vecty.Attribute("aria-valuemin", "0"),
					vecty.Attribute("aria-valuemax", "50"),
					vecty.Attribute("aria-valuenow", "20"),
					vecty.Attribute("aria-label", "Select Value"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mdc-slider__track-container"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("mdc-slider__track"),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mdc-slider__thumb-container"),
					),
					vecty.Tag(
						"svg",
						vecty.Markup(
							vecty.Class("mdc-slider__thumb"),
							vecty.Attribute("width", "21"),
							vecty.Attribute("height", "21"),
						),
						vecty.Tag(
							"circle",
							vecty.Markup(
								vecty.Attribute("cx", "10.5"),
								vecty.Attribute("cy", "10.5"),
								vecty.Attribute("r", "7.875"),
							),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("mdc-slider__focus-ring"),
						),
					),
				),
			),
		),
	)
}
