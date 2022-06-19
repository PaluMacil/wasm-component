package web

import "github.com/PaluMacil/wasm-component/wc"

// CopyrightFactory builds a TemplateRef of CopyrightComponent
type CopyrightFactory struct {
}

func (c CopyrightFactory) Register() {

}

// assert CopyrightFactory is a Registerer
var _ wc.Registerer = (*CopyrightFactory)(nil)

// CopyrightState hold state about the CopyrightComponent
type CopyrightState struct {
	Owners string
	Year   int
}

// CopyrightComponent holds the
type CopyrightComponent struct {
	//State wc.State[CopyrightState]
}
