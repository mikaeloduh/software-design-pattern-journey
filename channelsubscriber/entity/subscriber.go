package entity

// Subscriber type
type Subscriber struct {
	Name       string
	notifyCall func(c *Channel, v *Video, s *Subscriber)
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{Name: name}
}

func (s *Subscriber) OnNotify(c *Channel, v *Video) {
	s.notifyCall(c, v, s)
}

func (s *Subscriber) SetNotify(f func(c *Channel, v *Video, o *Subscriber)) {
	s.notifyCall = f
}
