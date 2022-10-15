package mahjong

type Player struct {
	Cards    []int
	Discards []int

	richi bool
}

func (p *Player) Clear() {
	p.Cards = make([]int, 0)
	p.Discards = make([]int, 0)
}

func (p *Player) CardsToString() (string, error) {
	return MahjongsToString(p.Cards)
}
