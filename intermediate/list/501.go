package list

func FindXTimes(teams string) bool {
	mapTeam := map[string]int{}
	for _, team := range teams {
		mapTeam[string(team)]++
	}
	intCnt := mapTeam[string(teams[0])]
	for _, val := range mapTeam {
		if intCnt != val {
			return false
		}
	}
	return true
}
