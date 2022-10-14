package test

import (
	"testing"

	. "github.com/konicyQWQ/mahjong/src/mahjong"
)

func TestMahjong(t *testing.T) {
	mahjong := Mahjong{}

	A := mahjong.NewPlayer()
	B := mahjong.NewPlayer()

	mahjong.InitDeck(114514)
	mahjong.InitPlayerCard()

	Acards, err := A.CardsToString()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	Bcards, err := B.CardsToString()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	t.Log("A: ", Acards)
	t.Log("B: ", Bcards)

	mahjong.PlayerDrawCard(A)

	Acards, err = A.CardsToString()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	err = mahjong.PlayerDiscard(A, "3m")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	Acards, err = A.CardsToString()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestDiscardNotExists(t *testing.T) {
	mahjong := Mahjong{}

	A := mahjong.NewPlayer()

	mahjong.InitDeck(114514)
	mahjong.InitPlayerCard()

	mahjong.PlayerDrawCard(A)

	Acards, err := A.CardsToString()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	t.Log("A: ", Acards)

	err = mahjong.PlayerDiscard(A, "3z")
	if err == nil {
		t.Log("Expect Error: ", err.Error())
		t.FailNow()
	}
}
