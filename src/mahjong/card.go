package mahjong

import (
	"errors"
	"fmt"
	"strings"
)

var CardSuffix = []string{"m", "s", "p", "z"}

func int2mahjong(i int) (string, error) {
	if i >= 34 || i < 0 {
		return "", errors.New(fmt.Sprintf("Invalid Card %d, which should be in [0, 34)", i))
	}
	return fmt.Sprintf("%d%s", i%9+1, CardSuffix[i/9]), nil
}

func mahjong2int(str string) (int, error) {
	var (
		i      int
		suffix string
	)
	n, err := fmt.Fscanf(strings.NewReader(str), "%d%s", &i, &suffix)
	if n != 2 || err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid string %s", str))
	}

	suffixIdx, err := find(CardSuffix, suffix)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid mahjong suffix %s", suffix))
	}

	return i - 1 + suffixIdx*9, nil
}
