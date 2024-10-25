package recursion

func FibonacciTail(n int32) int32 { // 末尾再帰を実装
	return fibonacciTailHelper(0, 1, n)
}

func fibonacciTailHelper(f1, f2, n int32) int32 {
	if n == 0 {
		return f1
	}
	return fibonacciTailHelper(f2, f1+f2, n-1)
}

// func Fibonacci(n int32) int32 { //再帰関数実装
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }
