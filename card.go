package main

import "fmt"

// How to write enums in Go
// See: http://blog.y-yuki.net/entry/2017/05/09/000000
type CardKind int

const (
	Diamond CardKind = iota
	Heart
	Clover
	Spade
)

func (kind CardKind) String() string {
	switch kind {
	case Diamond:
		return "ダイヤ"
	case Heart:
		return "ハート"
	case Clover:
		return "クラブ"
	case Spade:
		return "スペード"
	default:
		return "Unknown"
	}
}

type Card struct {
	number int
	kind CardKind
}

func (c Card) String() string {
	return fmt.Sprintf("%vの%d", c.kind, c.number)
}
