package libs

// CompositeExporter
type CompositeExporter struct {
	Children []IExporter
}

func NewCompositeExporter(children ...IExporter) *CompositeExporter {
	return &CompositeExporter{Children: children}
}

func (e *CompositeExporter) Write(s string) {
	for _, child := range e.Children {
		child.Write(s)
	}
}
