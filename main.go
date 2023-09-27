package main

import (
	"fmt"
)

type Multiply interface {
	Multiply(x, y []int) []int
}

//----------------------------------------------

type TraditionalMultiplication struct{}

func (s *TraditionalMultiplication) Multiply(num1, num2 []int) []int {
	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			temp := num1[i] * num2[j]
			p1, p2 := i+j, i+j+1
			sum := temp + result[p2]

			result[p1] += sum / 10
			result[p2] = sum % 10
		}
	}

	i := 0
	for i < len(result)-1 && result[i] == 0 {
		i++
	}

	return result[i:]
}

//----------------------------------------------

type KaratsubaMultiplication struct{}

func (k *KaratsubaMultiplication) Multiply(num1, num2 []int) []int {
	n := len(num1)

	if n <= 1 {
		result := make([]int, 2)
		result[0] = num1[0] * num2[0]
		if result[0] >= 10 {
			result[1] = result[0] / 10
			result[0] %= 10
		}
		return result
	}

	mid := n / 2

	a := num1[:mid]
	b := num1[mid:]
	c := num2[:mid]
	d := num2[mid:]

	ac := k.Multiply(a, c)
	bd := k.Multiply(b, d)

	ab := make([]int, n-mid)
	cd := make([]int, n-mid)
	for i := 0; i < len(ab); i++ {
		if i < len(a) {
			ab[i] = a[i]
		}
		if i < len(b) {
			ab[i] += b[i]
		}
		if i < len(c) {
			cd[i] = c[i]
		}
		if i < len(d) {
			cd[i] += d[i]
		}
	}
	abcd := k.Multiply(ab, cd)
	abcd = k.Subtract(abcd, ac)
	abcd = k.Subtract(abcd, bd)

	result := make([]int, 2*n)
	copy(result, ac)
	copy(result[mid:], abcd)
	copy(result[2*mid:], bd)

	i := 0
	for i < len(result)-1 && result[i] == 0 {
		i++
	}

	return result[i:]
}

func (k *KaratsubaMultiplication) Subtract(num1, num2 []int) []int {
	n := len(num1)
	result := make([]int, n)
	borrow := 0

	for i := n - 1; i >= 0; i-- {
		diff := num1[i] - borrow
		if i < len(num2) {
			diff -= num2[i]
		}

		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result[i] = diff
	}

	return result
}

//----------------------------------------------

type Calculate struct {
	strategy Multiply
}

func (c *Calculate) SetStrategy(strategy Multiply) {
	c.strategy = strategy
}

func (c *Calculate) Multiply(num1, num2 []int) []int {
	return c.strategy.Multiply(num1, num2)
}

func main() {
	task1 := &Calculate{}

	TraditionalMultiplication := &TraditionalMultiplication{}
	KaratsubaMultiplication := &KaratsubaMultiplication{}

	task1.SetStrategy(TraditionalMultiplication)

	x := []int{1, 2, 3, 4, 5}
	y := []int{5, 4, 3, 2, 1}

	result1 := task1.Multiply(x, y)

	fmt.Println("Результат умножения:")
	for _, digit := range result1 {
		fmt.Print(digit)
	}
	fmt.Println()

	task1.SetStrategy(KaratsubaMultiplication)
	result2 := task1.Multiply(x, y)

	fmt.Println("Результат умножения:")
	for _, digit := range result2 {
		fmt.Print(digit)
	}
	fmt.Println()
}
