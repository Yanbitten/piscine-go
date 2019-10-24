package piscine

func SortIntegerTable(table []int) {
	min_idx := 0
	for i := 0; i < len(table)-1; i++ {
		min_idx = i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[min_idx] {
				min_idx = j
			}
		}
		table[i], table[min_idx] = table[min_idx], table[i]
	}
}
