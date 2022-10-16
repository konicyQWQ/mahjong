package game

import (
	"fmt"
	"time"

	"github.com/konicyQWQ/mahjong/src/mahjong"
)

type GameIO interface {
	Read(int) (string, error)
	Write(int, string) error
}

func GameStart(io GameIO, playerNum int) {
	game := mahjong.Mahjong{}

	game.InitGame(int(time.Now().Unix()), 2)
	for i := 0; i < game.GetPlayerNumber(); i++ {
		ins := NewInsShowCard(*game.GetIthPlayer(i))
		io.Write(i, ins.ToJson())
	}

	for true {
		u := game.WhosTurn()
		player := game.GetIthPlayer(u)

		// 抽牌并检查卡组
		card, err := game.GetDeckFirstCard()
		if err != nil {
			// TODO: Game Over
			fmt.Println("No Card")
			return
		}

		// 更新玩家手牌
		player.DrawCard(*card)
		ins := NewInsShowCard(*player)
		io.Write(u, ins.ToJson())

		// 检查玩家操作指令
		t := Discard
		if player.CanHu(nil) {
			t = t | Hu
		}

		// 发送给该玩家可操作的指令
		ins = NewInsCanOperation((int)(t))
		io.Write(u, ins.ToJson())

		// 等待回复
		str, err := io.Read(u)
		if err != nil {
			fmt.Println("Read player operation failed")
			return
		}
		insPtr, err := NewInsFromJson(str)
		if err != nil {
			fmt.Println("Invalid Ins")
			return
		}

		// 根据玩家指令做不同操作
		// 丢牌
		if insPtr.Op == int(Discard) {
			// 玩家丢牌
			card, err := mahjong.NewCardFromString(insPtr.Extra)
			if err != nil {
				fmt.Println("Invalid Discard")
				return
			}
			player.Discard(*card)

			// 更新手牌
			ins = NewInsShowCard(*player)
			io.Write(u, ins.ToJson())

			// 更新别人看到的牌河
			for i := 0; i < game.GetPlayerNumber(); i++ {
				if i != u {
					ins = NewInsShowDisCard(u, *player)
					io.Write(i, ins.ToJson())
				}
			}

			game.NextTurn()
		}

		// 胡牌
		if insPtr.Op == Hu {
			// TODO: Game Over
			fmt.Printf("Game End, %d player win!\n", u)
			return
		}
	}
}
