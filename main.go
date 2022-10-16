package main

import (
	"errors"
	"fmt"

	"github.com/konicyQWQ/mahjong/src/game"
)

type IO struct{}

func (i IO) Read(u int) (string, error) {
	fmt.Printf("%d player, please input: ", u)

	var op string
	fmt.Scanf("%s", &op)

	if op == "d" {
		var card string
		fmt.Scanf("%s", &card)

		ins := game.Ins{
			Op:    int(game.Discard),
			Extra: card,
		}

		return ins.ToJson(), nil
	}

	return "", errors.New("No operation")
}

func (i IO) Write(u int, str string) error {
	ins, err := game.NewInsFromJson(str)
	if err != nil {
		fmt.Println("Invalid operation from server")
	}

	if ins.Type == game.ShowCard {
		fmt.Printf("%d player: %s\n", u, ins.Player.CardsToString())
	}

	if ins.Type == game.ShowDisCard {
		fmt.Printf("%s player discards: %s\n", ins.Extra, ins.Player.DisCardsToString())
	}

	if ins.Type == game.CanOperation {
		var str string
		if ins.Op&(int)(game.Discard) != 0 {
			str += " Discard"
		}
		if ins.Op&(int)(game.Hu) != 0 {
			str += " Hu"
		}

		fmt.Printf("%d player, your can: %s\n", u, str)
	}

	return nil
}

func main() {
	var io IO
	game.GameStart(io, 2)
}
