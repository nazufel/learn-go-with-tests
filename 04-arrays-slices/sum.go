package main

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(n ...[]int) (s []int) {

	var sums []int

	for _, numbers := range n {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(n ...[]int) []int {
	var sums []int

	for _, num := range n {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
