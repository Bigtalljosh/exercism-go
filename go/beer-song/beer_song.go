package beer

import (
	"errors"
	"fmt"
)

const (
	ZeroBottleVerse = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	OneBottleVerse  = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	TwoBottlesVerse = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	GenericVerse    = "%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n"
)

func Song() string {
	beerSong, _ := Verses(99, 0)
	return beerSong
}

func Verses(start, end int) (string, error) {
	if end > start {
		return "", errors.New("End cannot be greater than start!")
	}
	verses := ""
	for i := start; i >= end; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses += v + "\n"
	}
	return verses, nil
}

func Verse(bottles int) (string, error) {
	if bottles < 0 || bottles > 99 {
		return "", errors.New("You must enter a number of bottles between 0 and 99!")
	}

	switch {
	case bottles == 0:
		return ZeroBottleVerse, nil
	case bottles == 1:
		return OneBottleVerse, nil
	case bottles == 2:
		return TwoBottlesVerse, nil
	default:
		return fmt.Sprintf(GenericVerse, bottles, bottles, bottles-1), nil
	}
}
