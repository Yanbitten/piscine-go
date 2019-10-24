package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	answer := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		answer[i] = min + i
	}
	return answer
}
