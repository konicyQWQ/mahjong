package mahjong

import (
	"errors"
	"fmt"
)

func Find[T comparable](arr []T, el T) (int, error) {
	for i := range arr {
		if arr[i] == el {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprint(el, " Not Found In array"))
}
