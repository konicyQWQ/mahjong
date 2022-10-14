// 每个package 关键字之前，需要对包进行概述
package mahjong

import (
	"errors"
	"fmt"
	"math/rand"
)

type Mahjong struct {
	seed    int
	players []*Player
	deck    []int
}

func (m *Mahjong) InitDeck(seed int) {
	// init 1~9m, 1~9p, 1~9s, 1~7z
	m.deck = make([]int, 0, 136)
	for i := 0; i < 9+9+9+7; i++ {
		for j := 0; j < 4; j++ {
			m.deck = append(m.deck, i)
		}
	}

	// random shuffle
	m.seed = seed
	rand.Shuffle(len(m.deck), func(i, j int) {
		m.deck[i], m.deck[j] = m.deck[j], m.deck[i]
	})
}

func (m *Mahjong) InitPlayerCard() {
	for _, v := range m.players {
		for i := 1; i <= 13; i++ {
			m.PlayerDrawCard(v)
		}
	}
}

func (m *Mahjong) NewPlayer() *Player {
	p := &Player{}
	m.players = append(m.players, p)
	return p
}

func (m *Mahjong) PlayerDrawCard(p *Player) error {
	if m.IsGameEnd() {
		return errors.New("GameEnd, can not draw card")
	}

	p.cards = append(p.cards, m.deck[len(m.deck)-1])
	m.deck = append(m.deck[:len(m.deck)-1])
	return nil
}

func (m *Mahjong) PlayerDiscard(p *Player, card string) error {
	cardInt, err := mahjong2int(card)
	if err != nil {
		return err
	}

	idx, err := find(p.cards, cardInt)
	if err != nil {
		return errors.New(fmt.Sprintf("The %s is not in the player's cards", card))
	}

	p.discards = append(p.discards, cardInt)
	p.cards = append(p.cards[:idx], p.cards[idx+1:]...)
	return nil
}

func (m *Mahjong) IsGameEnd() bool {
	return len(m.deck) == 0
}
