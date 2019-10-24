package piscine

func AppendRange(min, max int) []int { //sdssg
	if min >= max {
		return nil
	}
	var answer []int
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
