package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, num := range numbers {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
