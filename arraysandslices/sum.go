package arraysandslices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	sums := make([]int, 0)
	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	sums := make([]int, 0)
	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
