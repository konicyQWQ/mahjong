package test

import (
	"testing"

	. "github.com/konicyQWQ/mahjong/src/mahjong"
)

func TestHu(t *testing.T) {
	var hu = []string{
		"11123345678883m",
		"23334445678992m",
		"11123444789996m",
		"11123456789991m",
		"11123456789992m",
		"11123456789993m",
		"11123456789994m",
		"11123456789995m",
		"11123456789996m",
		"11123456789997m",
		"11123456789998m",
		"11123456789999m",
		"123234m33s456p111z",
	}

	for _, v := range hu {
		cards, _ := String2Mahjongs(v)
		if CheckCardsCanHu(cards) != true {
			t.Fail()
		}
	}

}

func TestMahjong(t *testing.T) {
	mahjong := Mahjong{}

	A := mahjong.NewPlayer()
	B := mahjong.NewPlayer()

	mahjong.InitGame(114514)

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

	mahjong.InitGame(114514)

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
