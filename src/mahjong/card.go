package mahjong

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var CardSuffix = []string{"m", "s", "p", "z"}

func Int2mahjong(i int) (string, error) {
	if i >= 34 || i < 0 {
		return "", errors.New(fmt.Sprintf("Invalid Card %d, which should be in [0, 34)", i))
	}
	return fmt.Sprintf("%d%s", i%9+1, CardSuffix[i/9]), nil
}

func Mahjong2int(str string) (int, error) {
	var (
		i      int
		suffix string
	)
	n, err := fmt.Fscanf(strings.NewReader(str), "%d%s", &i, &suffix)
	if n != 2 || err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid string %s", str))
	}

	suffixIdx, err := Find(CardSuffix, suffix)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid mahjong suffix %s", suffix))
	}

	return i - 1 + suffixIdx*9, nil
}

func String2Mahjongs(str string) (number []int, err error) {
	for i, v := range CardSuffix {
		reg := regexp.MustCompile(fmt.Sprintf(`(\d+%s)?`, v))
		result := reg.FindAllString(str, -1)

		strArr := make([]string, 14)
		for _, matchStr := range result {
			if matchStr != "" {
				strArr = strings.Split(matchStr, v)
				break
			}
		}

		if len(strArr) == 2 {
			for _, v := range strArr[0] {
				card, err := strconv.Atoi(string([]rune{v}))
				if err != nil {
					return number, errors.New("Invalid string")
				}

				number = append(number, card-1+i*9)
			}
		}
	}

	return number, nil
}

func MahjongsToString(cards []int) (string, error) {
	len := len(cards)
	if len == 0 {
		return "", nil
	}

	sort.Sort(sort.IntSlice(cards[:len-1]))

	str := ""
	for idx, v := range cards {
		if idx >= len-2 || cards[idx+1]/9 != v/9 {
			mahjong, err := Int2mahjong(v)
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
