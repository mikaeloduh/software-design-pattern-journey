package entity

// IHandler interface
type IHandler interface {
	Handle(c1Ptr, c2Ptr *Sprite)
}

// TemplateHandler is it
type TemplateHandler struct {
	Match      func(c1Ptr, c2Ptr *Sprite) bool
	DoHandling func(c1Ptr, c2Ptr *Sprite)
	Next       IHandler
}

func (h TemplateHandler) Handle(c1Ptr, c2Ptr *Sprite) {
	if h.Match(c1Ptr, c2Ptr) {
		h.DoHandling(c1Ptr, c2Ptr)
	} else if h.Next != nil {
		h.Next.Handle(c1Ptr, c2Ptr)
	}
}
