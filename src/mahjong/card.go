package mahjong

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Card struct {
	Number   int
	CardType CardType
}

func (c *Card) ToString() string {
	return fmt.Sprintf("%d%s", c.Number, CardType2str(c.CardType))
}

func (c *Card) ToInt() int {
	return c.Number - 1 + int(c.CardType)*9
}

func NewCardFromString(str string) (*Card, error) {
	var (
		i      int
		suffix string
	)
	n, err := fmt.Fscanf(strings.NewReader(str), "%d%s", &i, &suffix)
	if n != 2 || err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid string %s", str))
	}

	suffixIdx, err := Find(cardTypeChar[:], suffix)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid mahjong suffix %s", suffix))
	}

	card := Card{
		Number:   i,
		CardType: (CardType)(suffixIdx),
	}
	return &card, nil
}

func NewCardFromInt(i int) (*Card, error) {
	if i >= 34 || i < 0 {
		return nil, errors.New(fmt.Sprintf("Invalid Card %d, which should be in [0, 34)", i))
	}

	n := i%9 + 1
	suffixIdx := i / 9

	card := Card{
		Number:   n,
		CardType: (CardType)(suffixIdx),
	}
	return &card, nil
}

func Cards2Ints(cards []Card) (ret []int) {
	for _, c := range cards {
		ret = append(ret, c.ToInt())
	}

	return ret
}

func CardsToString(cards []Card) (string, error) {
	len := len(cards)
	if len == 0 {
		return "", nil
	}

	cardInts := Cards2Ints(cards)
	sort.Sort(sort.IntSlice(cardInts[:len-1]))

	str := ""
	for idx, v := range cardInts {
		if idx >= len-2 || cardInts[idx+1]/9 != v/9 {
			card, err := NewCardFromInt(v)
			if err != nil {
				return "", err
			}
			mahjong := card.ToString()
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
