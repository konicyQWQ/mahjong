package mahjong

import (
	"fmt"
	"sort"
)

type Player struct {
	cards    []int
	discards []int
}

func (p *Player) CardsToString() (string, error) {
	len := len(p.cards)
	if len == 0 {
		return "", nil
	}

	sort.Sort(sort.IntSlice(p.cards[:len-1]))

	str := ""
	for idx, v := range p.cards {
		if idx >= len-2 || p.cards[idx+1]/9 != v/9 {
			mahjong, err := int2mahjong(v)
			if err != nil {
				return "", err
			}
			if idx == len-1 {
				str += "   "
			}
			str += mahjong
		} else {
			str += fmt.Sprintf("%d", v%9+1)
		}
	}

	return str, nil
}
