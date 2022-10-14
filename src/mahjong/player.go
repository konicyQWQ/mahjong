package mahjong

type Player struct {
	cards    []int
	discards []int
}

func (p *Player) Clear() {
	p.cards = make([]int, 0)
	p.discards = make([]int, 0)
}

func (p *Player) CardsToString() (string, error) {
	return MahjongsToString(p.cards)
}
