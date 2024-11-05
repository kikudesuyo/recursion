package list

import (
	"errors"
	"strconv"
)

func cardNum(cVal string) (int, error) {
	switch cVal {
	case "A":
		return 1, nil
	case "J":
		return 11, nil
	case "Q":
		return 12, nil
	case "K":
		return 13, nil
	}
	num, err := strconv.Atoi(cVal)
	if err != nil {
		return 0, errors.New("invalid card value")
	}
	return num, nil
}

func WinnerBlackjack(playerCards []string, houseCards []string) bool {
	playerNum := 0
	for _, card := range playerCards {
		cVal := string([]rune(card)[1:])
		num, err := cardNum(cVal)
		if err != nil {
			return false
		}
		playerNum += num
	}
	houseNum := 0
	for _, card := range houseCards {
		cVal := string([]rune(card)[1:])
		num, err := cardNum(cVal)
		if err != nil {
			return false
		}
		houseNum += num
	}
	if playerNum > 21 {
		return false
	}
	if houseNum > 21 {
		return true
	}
	return playerNum > houseNum
}
