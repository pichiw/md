package md

import (
	"strconv"
	"sync"

	"github.com/gowasm/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func NewSlider(opts SliderOptions) *Slider {
	return &Slider{
		opts: opts,
	}
}

type OnSliderChange func(s *Slider)

type SliderOptions struct {
	Min   float64
	Max   float64
	Value float64

	OnChange OnSliderChange
}
type Slider struct {
	vecty.Core

	opts      SliderOptions
	optsMutex sync.Mutex
}

func (s *Slider) Value() float64 {
	s.optsMutex.Lock()
	defer s.optsMutex.Unlock()
	return s.opts.Value
}

func (s *Slider) Mount() {
	doc := js.Global().Get("document")

	root := doc.Call("getElementById", "slider-example")
	sliderEl := root.Call("querySelector", "#continuous-mdc-slider")
	slider := js.Global().Get("mdc").Get("slider").Get("MDCSlider").New(sliderEl)
	slider.Call("listen", "MDCSlider:change",
		js.NewCallback(
			func(vs []js.Value) {
				s.optsMutex.Lock()
				s.opts.Value = slider.Get("value").Float()
				s.optsMutex.Unlock()

				oc := s.opts.OnChange
				if oc == nil {
					return
				}
				oc(s)
			},
		),
	)
}

func (s *Slider) Render() vecty.ComponentOrHTML {
	minStr := strconv.FormatFloat(s.opts.Min, 'f', -1, 64)
	maxStr := strconv.FormatFloat(s.opts.Max, 'f', -1, 64)
	valueStr := strconv.FormatFloat(s.opts.Value, 'f', -1, 64)
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
					vecty.Attribute("aria-valuemin", minStr),
					vecty.Attribute("aria-valuemax", maxStr),
					vecty.Attribute("aria-valuenow", valueStr),
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
