package entity

import (
	"collisiongame/commons"
)

/**
 * | Subject \ target |  Hero  | Water  |  Fire  |
 * | :--------------: | :----: | :----: | :----: |
 * |       Hero       |  fail  |  +10   |  -10   |
 * |      Water       | remove |  fail  | remove |
 * |       Fire       | remove | remove |  fail  |
 */

type HeroHeroHandler struct {
	// Hero -> Hero
}

func NewHeroHeroHandler(next IHandler) *TemplateHandler {
	h := HeroHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Hero{}) && commons.IsSameType(*c2Ptr, &Hero{})
}

func (h HeroHeroHandler) DoHandling(_, _ *Sprite) {
	return
}

type HeroWaterHandler struct {
	// Hero -> Water
}

func NewHeroWaterHandler(next IHandler) *TemplateHandler {
	h := HeroWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Hero{}) && commons.IsSameType(*c2Ptr, &Water{})
}

func (h HeroWaterHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	(*c1Ptr).(*Hero).SetHp(+10)
	*c2Ptr = nil
	*c2Ptr = *c1Ptr
	*c1Ptr = nil
}

type HeroFireHandler struct {
	// Hero -> Fire
}

func NewHeroFireHandler(next IHandler) *TemplateHandler {
	h := HeroFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h HeroFireHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Hero{}) && commons.IsSameType(*c2Ptr, &Fire{})
}

func (h HeroFireHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	(*c1Ptr).(*Hero).SetHp(-10)
	*c2Ptr = nil
	*c2Ptr = *c1Ptr
	*c1Ptr = nil
}

type WaterHeroHandler struct {
	// Water -> Hero
}

func NewWaterHeroHandler(next IHandler) *TemplateHandler {
	h := WaterHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Water{}) && commons.IsSameType(*c2Ptr, &Hero{})
}

func (h WaterHeroHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	(*c2Ptr).(*Hero).SetHp(+10)
	*c1Ptr = nil
}

type WaterWaterHandler struct {
	// Water -> Water
}

func NewWaterWaterHandler(next IHandler) *TemplateHandler {
	h := WaterWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Water{}) && commons.IsSameType(*c2Ptr, &Water{})
}

func (h WaterWaterHandler) DoHandling(_, _ *Sprite) {
	return
}

type WaterFireHandler struct {
	// Water -> Fire
}

func NewWaterFireHandler(next IHandler) *TemplateHandler {
	h := WaterFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h WaterFireHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Water{}) && commons.IsSameType(*c2Ptr, &Fire{})
}

func (h WaterFireHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	*c1Ptr = nil
	*c2Ptr = nil
}

type FireHeroHandler struct {
	// Fire -> Hero
}

func NewFireHeroHandler(next IHandler) *TemplateHandler {
	h := FireHeroHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireHeroHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Fire{}) && commons.IsSameType(*c2Ptr, &Hero{})
}

func (h FireHeroHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	*c1Ptr = nil
	(*c2Ptr).(*Hero).SetHp(-10)
}

type FireWaterHandler struct {
	// Fire -> Water
}

func NewFireWaterHandler(next IHandler) *TemplateHandler {
	h := FireWaterHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireWaterHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Fire{}) && commons.IsSameType(*c2Ptr, &Water{})
}

func (h FireWaterHandler) DoHandling(c1Ptr, c2Ptr *Sprite) {
	*c1Ptr = nil
	*c2Ptr = nil
}

type FireFireHandler struct {
	// Fire -> Fire
}

func NewFireFireHandler(next IHandler) *TemplateHandler {
	h := FireFireHandler{}
	return &TemplateHandler{
		Match:      h.Match,
		DoHandling: h.DoHandling,
		Next:       next,
	}
}

func (h FireFireHandler) Match(c1Ptr, c2Ptr *Sprite) bool {
	return commons.IsSameType(*c1Ptr, &Fire{}) && commons.IsSameType(*c2Ptr, &Fire{})
}

func (h FireFireHandler) DoHandling(_, _ *Sprite) {
	return
}
