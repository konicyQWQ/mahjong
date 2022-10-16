package mahjong

import (
	"errors"
	"math/rand"
)

type Mahjong struct {
	seed    int
	players []Player
	deck    []Card
	turn    int
}

func (m *Mahjong) initSeed(seed int) {
	m.seed = seed
	rand.Seed((int64)(seed))
}

func (m *Mahjong) initDeck() {
	m.deck = make([]Card, 0, 136)
	for i := 0; i < 9+9+9+7; i++ {
		for j := 0; j < 4; j++ {
			v, _ := NewCardFromInt(i)
			m.deck = append(m.deck, *v)
		}
	}

	rand.Shuffle(len(m.deck), func(i, j int) {
		m.deck[i], m.deck[j] = m.deck[j], m.deck[i]
	})
}

func (m *Mahjong) initPlayer(playerNum int) {
	for i := 0; i < playerNum; i++ {
		m.players = append(m.players, Player{})
	}
}

func (m *Mahjong) initPlayerCard() {
	for j := 0; j < 13; j++ {
		for i := 0; i < len(m.players); i++ {
			v, err := m.GetDeckFirstCard()
			if err != nil {
				panic("initPlayerCard Internal error")
			}
			m.players[i].DrawCard(*v)
		}
	}
}

func (m *Mahjong) GetDeckFirstCard() (*Card, error) {
	if len(m.deck) == 0 {
		return nil, errors.New("No Card")
	}

	ret := m.deck[len(m.deck)-1]
	m.deck = append(m.deck[:len(m.deck)-1])

	return &ret, nil
}

func (m *Mahjong) InitGame(seed int, playerNum int) {
	m.initSeed(seed)
	m.initDeck()
	m.initPlayer(playerNum)
	m.initPlayerCard()
}

func (m *Mahjong) GetPlayerNumber() int {
	return len(m.players)
}

func (m *Mahjong) GetIthPlayer(i int) *Player {
	return &m.players[i]
}

func (m *Mahjong) WhosTurn() int {
	return m.turn
}

func (m *Mahjong) NextTurn() {
	m.turn = (m.turn + 1) % m.GetPlayerNumber()
}
