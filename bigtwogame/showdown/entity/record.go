package entity

import "bigtwogame/template"

type Result struct {
	Player template.IPlayer[ShowDownCard]
	Card   ShowDownCard
	Win    bool
}

type RoundResult []Result

type Record []RoundResult
