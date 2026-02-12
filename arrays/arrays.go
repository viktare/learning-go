package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, n := range numsToSum {
		sums = append(sums, Sum(n))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
