package game

import (
	"encoding/json"
	"fmt"

	"github.com/konicyQWQ/mahjong/src/mahjong"
)

type InsType int
type OperationType int

const (
	ShowCard InsType = iota + 0
	ShowDisCard
	CanOperation
	Operation
)

const (
	Discard OperationType = 1
	Pong                  = 2
	Chi                   = 4
	Hu                    = 8
)

type Ins struct {
	Type   InsType
	Player mahjong.Player
	Op     int
	Extra  string
}

func (i *Ins) ToJson() string {
	v, _ := json.Marshal(i)
	return string(v)
}

func NewInsShowCard(player mahjong.Player) Ins {
	ins := Ins{
		Type:   ShowCard,
		Player: player,
	}
	return ins
}

func NewInsShowDisCard(u int, player mahjong.Player) Ins {
	ins := Ins{
		Type:   ShowDisCard,
		Player: player,
		Extra:  fmt.Sprintf("%d", u),
	}
	return ins
}

func NewInsCanOperation(t int) Ins {
	ins := Ins{
		Type: CanOperation,
		Op:   t,
	}
	return ins
}

func NewInsFromJson(str string) (*Ins, error) {
	var ins Ins

	err := json.Unmarshal([]byte(str), &ins)
	if err != nil {
		return nil, err
	}

	return &ins, nil
}
