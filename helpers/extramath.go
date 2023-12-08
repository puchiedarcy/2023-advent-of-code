package helpers

import "sort"

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func FindGCD(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = gcd(numbers[i], result)

		if result == 1 {
			return 1
		}
	}
	return result
}

func lcm(a int, b int, gcd int) int {
	return a * b / gcd
}

func FindLCM(numbers []int) int {
	sort.Ints(numbers)
	result := numbers[0]
	gcd := FindGCD(numbers)
	for i := 1; i < len(numbers); i++ {
		result = lcm(numbers[i], result, gcd)
	}
	return result
}
