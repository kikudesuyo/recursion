package controlflow

func DoYouFail(s string) string {
	cnt := 0
	for _, char := range s {
		if string(char) == "A" {
			cnt++
		}
		if cnt >= 3 {
			return "fail"
		}
	}
	return "pass"
}
