package recursion

import "strconv"

// 再帰コード
func splitAndAdd(digits int64) int32 {
	var sum int32 = 0
	strDigits := strconv.FormatInt(digits, 10)
	for _, char := range strDigits {
		sum += int32(char - '0')
	}
	return sum
}

// // 普通のコード

func RecursiveDigitsAdded(digits int64) int32 {
	// 最初の数字の各桁を足す
	result := splitAndAdd(digits)
	total := result

	// resultが2桁以上の間、処理を続ける
	for result >= 10 {
		result = splitAndAdd(int64(result))
		total += result
	}

	return total
}
