package main

import (
	"fmt"
)

var num2en map[int]string
var post2en map[int]string

func ss(x int) string {
	var r string
	if x >= 100 {
		r = " " + num2en[x/100] + " Hundred"
		x = x % 100
	}
	if x >= 20 {
		r += " " + num2en[x/10*10]
		x = x % 10
	}
	if x > 0 {
		r += " " + num2en[x]
	}
	return r
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%1000)
		num /= 1000
	}

	var r string
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > 0 {
			r += ss(nums[i])
			if i > 0 {
				r += " " + post2en[i]
			}
		}
	}
	return r[1:]
}

func init() {
	num2en = map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}

	post2en = map[int]string{
		0: "",
		1: "Thousand",
		2: "Million",
		3: "Billion",
	}
}

func main() {
	r := numberToWords(100001)
	fmt.Println(r)
}
