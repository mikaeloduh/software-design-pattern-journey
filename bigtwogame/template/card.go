package template

type ICard interface {
	String() string
}

type IPattern []ICard
