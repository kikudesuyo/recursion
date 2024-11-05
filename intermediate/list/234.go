package list

func CharInBagOfWordsCount(bagOfWords []string, keyCharacter byte) int32 {
	var count int32
	for _, word := range bagOfWords {
		for i := 0; i < len(word); i++ {
			if word[i] == keyCharacter {
				count++
			}
		}
	}
	return count

}
