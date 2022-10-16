package mahjong

type CardType int

const (
	M CardType = iota + 1
	P
	S
	Z
)

var cardTypeChar = [4]string{"m", "p", "s", "z"}

func CardType2str(ct CardType) string {
	return cardTypeChar[ct]
}
