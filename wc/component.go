package wc

type Stateful interface {
	any
}

// TODO: a component should render. The generic
//       type should probably be the component
//       and a SpecificComponent could then embed
//       a Component[SpecificState].

type ComponentFactory func(templateRef *TemplateRef) error

type Component interface {
	Render() error
}

// Registerer registers a factory in the template registry
type Registerer interface {
	Register()
}

type TemplateRefFactory struct {
	templateFetcher TemplateFetcher
}
