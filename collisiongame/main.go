package main

import (
	"math/rand"
	"reflect"
)

// Sprite and its friends
type Sprite interface {
	String() string
}

func RandNewSprite() Sprite {
	return [3]func() Sprite{
		func() Sprite { return NewHero() },
		func() Sprite { return NewWater() },
		func() Sprite { return NewFire() },
	}[rand.Intn(3)]()
}

type Hero struct {
	Sprite
	hp int
}

func (h *Hero) String() string {
	return "H"
}

func (h *Hero) SetHp(n int) {
	h.hp += n
}

func NewHero() *Hero {
	return &Hero{hp: 30}
}

type Water struct{}

func NewWater() *Water {
	return &Water{}
}

func (w *Water) String() string {
	return "W"
}

type Fire struct{}

func NewFire() *Fire {
	return &Fire{}
}

func (f *Fire) String() string {
	return "F"
}

// World the happy sprites world
type World struct {
	coord   [30]Sprite
	handler IHandler
}

func NewWorld(h IHandler) *World {
	w := &World{handler: h}
	w.Init()
	return w
}

func (w *World) Init() {
	numbers := make([]int, 30)
	for i := 0; i < 30; i++ {
		numbers[i] = i
	}
	for i := len(numbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	for _, num := range numbers[:10] {
		w.coord[num] = RandNewSprite()
	}
}

// IHandler interface
type IHandler interface {
	Handle(c1, c2 *Sprite)
}

// TemplateHandler is it
type TemplateHandler struct {
	Match      func(c1, c2 *Sprite) bool
	DoHandling func(c1, c2 *Sprite)
	Next       IHandler
}

func (h TemplateHandler) Handle(c1, c2 *Sprite) {
	if h.Match(c1, c2) {
		h.DoHandling(c1, c2)
	} else if h.Next != nil {
		h.Next.Handle(c1, c2)
	}
}

type HeroHeroHandler struct{}

func NewHeroHeroHandler(next IHandler) *TemplateHandler {
	h := HeroHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Hero{}) && isSameType(*c2Ptr, &Hero{})
}

func (h HeroHeroHandler) DoHandling(_, _ *Sprite) {
	// Hero -> Hero
	return
}

type HeroWaterHandler struct{}

func NewHeroWaterHandler(next IHandler) *TemplateHandler {
	h := HeroWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Hero{}) && isSameType(*c2Ptr, &Water{})
}

func (h HeroWaterHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Hero -> Water
	(*c1Ptr).(*Hero).SetHp(+10)
	*c2Ptr = nil
	*c2Ptr = *c1Ptr
	*c1Ptr = nil
}

type HeroFireHandler struct{}

func NewHeroFireHandler(next IHandler) *TemplateHandler {
	h := HeroFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroFireHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Hero{}) && isSameType(*c2Ptr, &Fire{})
}

func (h HeroFireHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Hero -> Fire
	(*c1Ptr).(*Hero).SetHp(-10)
	*c2Ptr = nil
	*c2Ptr = *c1Ptr
	*c1Ptr = nil
}

type WaterHeroHandler struct{}

func NewWaterHeroHandler(next IHandler) *TemplateHandler {
	h := WaterHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Water{}) && isSameType(*c2Ptr, &Hero{})
}

func (h WaterHeroHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Water -> Hero
	(*c2Ptr).(*Hero).SetHp(+10)
	*c1Ptr = nil
}

type WaterWaterHandler struct{}

func NewWaterWaterHandler(next IHandler) *TemplateHandler {
	h := WaterWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Water{}) && isSameType(*c2Ptr, &Water{})
}

func (h WaterWaterHandler) DoHandling(_, _ *Sprite) {
	// Water -> Water
	return
}

type WaterFireHandler struct{}

func NewWaterFireHandler(next IHandler) *TemplateHandler {
	h := WaterFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterFireHandler) Match(c1, c2 *Sprite) bool {
	return isSameType(*c1, &Water{}) && isSameType(*c2, &Fire{})
}

func (h WaterFireHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Water -> Fire
	*c1Ptr = nil
	*c2Ptr = nil
}

type FireHeroHandler struct{}

func NewFireHeroHandler(next IHandler) *TemplateHandler {
	h := FireHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Fire{}) && isSameType(*c2Ptr, &Hero{})
}

func (h FireHeroHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Fire -> Hero
	*c1Ptr = nil
	(*c2Ptr).(*Hero).SetHp(-10)
}

type FireWaterHandler struct{}

func NewFireWaterHandler(next IHandler) *TemplateHandler {
	h := FireWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Fire{}) && isSameType(*c2Ptr, &Water{})
}

func (h FireWaterHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	// Fire -> Water
	*c1Ptr = nil
	*c2Ptr = nil
}

type FireFireHandler struct{}

func NewFireFireHandler(next IHandler) *TemplateHandler {
	h := FireFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireFireHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return isSameType(*c1Ptr, &Fire{}) && isSameType(*c2Ptr, &Fire{})
}

func (h FireFireHandler) DoHandling(_, _ *Sprite) {
	return
}

/**
 * | Subject \ target |  Hero  | Water  |  Fire  |
 * | :--------------: | :----: | :----: | :----: |
 * |       Hero       |  fail  |  +10   |  -10   |
 * |      Water       | remove |  fail  | remove |
 * |       Fire       | remove | remove |  fail  |
 */

func (w *World) Move(x1 int, x2 int) {
	// TODO: isValidMove(x1, x2)

	c1Ptr := &w.coord[x1]
	c2Ptr := &w.coord[x2]

	// toCollide and move
	w.handler.Handle(c1Ptr, c2Ptr)
}

func isSameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func main() {
	println("hello world")

	h := NewHeroHeroHandler(NewHeroWaterHandler(NewHeroFireHandler(NewWaterHeroHandler(NewWaterWaterHandler(NewWaterFireHandler(NewFireHeroHandler(NewFireWaterHandler(NewFireFireHandler(nil)))))))))
	_ = NewWorld(h)
}
