package recursion

func HowLongToReachFundGoal(capitalMoney int32, goalMoney int32, interest int32) int32 {
	if capitalMoney >= goalMoney {
		return 0
	}
	currentCapitalMoney := float64(capitalMoney)
	currentgoalMoney := float64(goalMoney)
	for currentYear := 0; currentYear <= 80; currentYear++ {
		currentCapitalMoney *= float64(float64(interest)/100 + 1)
		if currentYear%2 == 0 {
			currentgoalMoney *= 1.03
		} else {
			currentgoalMoney *= 1.02
		}
		if currentCapitalMoney >= currentgoalMoney {
			return int32(currentYear)
		}
	}
	return 80
}
