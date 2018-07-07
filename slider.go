package md

import (
	"strconv"
	"sync"

	"github.com/gowasm/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func NewSlider(opts SliderOptions) *Slider {
	if opts.Step == 0 {
		opts.Step = 1
	}
	return &Slider{
		divID: RandID(),
		opts:  opts,
	}
}

type OnSliderChange func(s *Slider)
type OnSliderInput func(s *Slider)

type SliderOptions struct {
	Min   float64
	Max   float64
	Value float64
	Step  float64

	OnChange OnSliderChange
	OnInput  OnSliderInput
}
type Slider struct {
	vecty.Core

	divID string

	slider     js.Value
	sliderOnce sync.Once

	onChangeCB js.Callback
	onInputCB  js.Callback

	opts      SliderOptions
	optsMutex sync.Mutex
}

func (s *Slider) Value() float64 {
	s.optsMutex.Lock()
	defer s.optsMutex.Unlock()
	return s.opts.Value
}

func (s *Slider) onChange(vs []js.Value) {
	s.optsMutex.Lock()
	s.opts.Value = s.slider.Get("value").Float()
	s.optsMutex.Unlock()

	oc := s.opts.OnChange
	if oc == nil {
		return
	}
	oc(s)
}

func (s *Slider) onInput(vs []js.Value) {
	s.optsMutex.Lock()
	s.opts.Value = s.slider.Get("value").Float()
	s.optsMutex.Unlock()

	oi := s.opts.OnInput
	if oi == nil {
		return
	}
	oi(s)
}

func (s *Slider) Mount() {
	s.sliderOnce.Do(func() {
		doc := js.Global().Get("document")

		sliderEl := doc.Call("getElementById", s.divID)
		s.slider = js.Global().Get("mdc").Get("slider").Get("MDCSlider").New(sliderEl)

		s.onChangeCB = js.NewCallback(s.onChange)
		s.onInputCB = js.NewCallback(s.onInput)

		s.slider.Call("listen", "MDCSlider:change", s.onChangeCB)
		s.slider.Call("listen", "MDCSlider:input", s.onInputCB)
	})
}

func (s *Slider) Unmount() {
	s.onChangeCB.Release()
	s.onInputCB.Release()
}

func (s *Slider) Render() vecty.ComponentOrHTML {
	minStr := strconv.FormatFloat(s.opts.Min, 'f', -1, 64)
	maxStr := strconv.FormatFloat(s.opts.Max, 'f', -1, 64)
	valueStr := strconv.FormatFloat(s.opts.Value, 'f', -1, 64)
	stepStr := strconv.FormatFloat(s.opts.Step, 'f', -1, 64)
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-slider", "mdc-slider--discrete"),
			vecty.Attribute("id", s.divID),
			vecty.Attribute("tabindex", "0"),
			vecty.Attribute("role", "slider"),
			vecty.Attribute("aria-valuemin", minStr),
			vecty.Attribute("aria-valuemax", maxStr),
			vecty.Attribute("aria-valuenow", valueStr),
			vecty.Attribute("data-step", stepStr),
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
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-slider__pin"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("mdc-slider__pin-value-marker"),
					),
				),
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
	)
}
