package mahjong

import (
	"encoding/json"
	"errors"
)

type PongChi struct {
	Cards []Card
}

type Player struct {
	Cards    []Card
	Discards []Card
	PongChi  []PongChi
}

func (p *Player) DrawCard(c Card) {
	p.Cards = append(p.Cards, c)
}

func (p *Player) Discard(c Card) error {
	idx, err := Find(p.Cards, c)
	if err != nil {
		return errors.New("Must Discard Your Card")
	}

	p.Cards = append(p.Cards[:idx], p.Cards[idx+1:]...)
	p.Discards = append(p.Discards, c)

	return nil
}

func (p *Player) GetAllCards() (ret []Card) {
	ret = append(ret, p.Cards...)
	for _, v := range p.PongChi {
		ret = append(ret, v.Cards...)
	}
	return ret
}

func (p *Player) CanHu(c *Card) bool {
	cards := p.GetAllCards()
	if c != nil {
		cards = append(cards, *c)
	}

	return CheckCardsCanHu(Cards2Ints(cards))
}

func (p *Player) ToJson() string {
	arr, _ := json.Marshal(p)
	return string(arr)
}

func (p *Player) CardsToString() string {
	str, _ := CardsToString(p.Cards)
	return str
}

func (p *Player) DisCardsToString() string {
	var str string
	for _, v := range p.Discards {
		str += v.ToString() + " "
	}
	return str
}
