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
	if len(num1) < len(num2) {
		newNum := make([]int, len(num2))
		num1 = k.Summa(newNum, num1)
	}

	if len(num2) < len(num1) {
		newNum := make([]int, len(num1))
		num2 = k.Summa(newNum, num2)
	}

	n := len(num1)

	if n <= 1 {
		result := make([]int, 2)
		result[1] = num1[0] * num2[0]
		if result[1] >= 10 {
			result[0] = result[1] / 10
			result[1] %= 10
		}

		i := 0
		for i < len(result)-1 && result[i] == 0 {
			i++
		}
		return result[i:]
	}

	mid := n / 2
	mid1 := mid
	if n%2 != 0 {
		mid1++
	}

	a := num1[:mid1]
	b := num1[mid1:]
	c := num2[:mid1]
	d := num2[mid1:]

	ac := k.Multiply(a, c)
	bd := k.Multiply(b, d)

	ab := k.Summa(a, b)
	cd := k.Summa(c, d)
	abcd := k.Multiply(ab, cd)
	abcd = k.Subtract(abcd, ac)
	abcd = k.Subtract(abcd, bd)

	sum1 := make([]int, mid*2+len(ac))
	acNum := 0
	for i := 0; i < len(sum1); i++ {
		if acNum == len(ac) {
			break
		}
		sum1[i] = ac[i]
		acNum++
	}

	sum2 := make([]int, mid+len(abcd))
	abcdNum := 0
	for i := 0; i < len(sum2); i++ {
		if abcdNum == len(abcd) {
			break
		}
		sum2[i] = abcd[i]
		abcdNum++
	}

	result := make([]int, 2*n)
	result = k.Summa(result, sum1)
	result = k.Summa(result, sum2)
	result = k.Summa(result, bd)

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

		j := i - (len(num1) - len(num2))
		if j >= 0 {
			diff -= num2[j]
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

func (k *KaratsubaMultiplication) Summa(num1 []int, num2 []int) []int {
	maxLen := len(num1)
	if len(num2) > maxLen {
		maxLen = len(num2)
	}

	result := make([]int, maxLen+1)
	carry := 0

	for i := 0; i < maxLen; i++ {
		sum := carry
		if i < len(num1) {
			sum += num1[len(num1)-1-i]
		}
		if i < len(num2) {
			sum += num2[len(num2)-1-i]
		}

		result[maxLen-i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		result[0] = carry
	} else {
		result = result[1:]
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
	y := []int{0}

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
