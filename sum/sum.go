package main

func Sum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, n := range numbersToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			tail := n[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
