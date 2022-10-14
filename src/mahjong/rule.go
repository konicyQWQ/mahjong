package mahjong

import (
	"errors"
	"sort"
)

func CheckCardsCanHu(cards []int) bool {
	if len(cards) != 14 {
		return false
	}

	tmpCards := make([]int, 14)
	copy(tmpCards, cards)

	return __checkCardsCanHu(tmpCards)
}

func __checkCardsCanHu(cards []int) bool {
	sort.Sort(sort.IntSlice(cards))
	used := make([]bool, 14)

	for idx, v := range cards[:13] {
		if v == cards[idx+1] {
			used[idx] = true
			used[idx+1] = true

			if __checkCardsCanHu_Dfs(cards, used) {
				return true
			}

			used[idx] = false
			used[idx+1] = false
		}
	}

	return false
}

func __checkCardsCanHu_Dfs(cards []int, used []bool) (ok bool) {

	findNoUsed := func(el int) (int, error) {
		for i, v := range cards {
			if used[i] == false && v == el {
				return i, nil
			}
		}
		return -1, errors.New("not find")
	}

	// 检查是否所有手牌用上了
	ok = true
	for _, v := range used {
		if v == false {
			ok = false
			break
		}
	}
	if ok {
		return true
	}

	for i, v := range cards {
		if used[i] == false {
			// 检查刻子
			used[i] = true

			idx, err := findNoUsed(v)
			if err == nil {
				used[idx] = true
				idx2, err2 := findNoUsed(v)
				if err2 == nil {
					used[idx2] = true
					ok = ok || __checkCardsCanHu_Dfs(cards, used)
					used[idx2] = false
				}
				used[idx] = false
			}

			used[i] = false

			// 检查顺子
			if v%9+1 <= 7 {
				used[i] = true

				idx, err := findNoUsed(v + 1)
				if err == nil {
					used[idx] = true
					idx2, err2 := findNoUsed(v + 2)
					if err2 == nil {
						used[idx2] = true
						ok = ok || __checkCardsCanHu_Dfs(cards, used)
						used[idx2] = false
					}
					used[idx] = false
				}

				used[i] = false
			}

			break
		}
	}

	return ok
}
