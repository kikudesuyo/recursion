package list

func HasSameType(user1 string, user2 string) bool {
	if len(user1) != len(user2) {
		return false
	}
	hashMap := make(map[string]string)
	for i := 0; i < len(user1); i++ {
		if _, ok := hashMap[string(user1[i])]; ok {
			if hashMap[string(user1[i])] != string(user2[i]) {
				return false
			}
		} else {
			hashMap[string(user1[i])] = string(user2[i])
		}
	}
	vals := make(map[string]bool)
	for _, v := range hashMap {
		vals[v] = true
	}
	return len(vals) == len(hashMap)
}
