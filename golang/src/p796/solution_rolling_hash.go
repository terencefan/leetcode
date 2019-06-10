package main

import (
	"fmt"
)

const BASE = 997
const MOD = 2<<62 - 1

func hash(A string) (h int64) {
	for i := range A {
		h *= BASE
		h += int64(A[i])
		h %= MOD
	}
	return
}

func compare(A, B string, i int) bool {
	for k := 0; k < len(B); k++ {
		if A[i] != B[k] {
			return false
		}
		i++
		i %= len(A)
	}
	return true
}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	} else if A == B {
		return true
	}

	length := len(A)
	rollingBase := int64(1)
	for i := 0; i < length; i++ {
		rollingBase *= BASE
		rollingBase %= MOD
	}

	hashA, hashB := hash(A), hash(B)

	for i := 0; i < length; i++ {
		if hashA == hashB {
			if compare(A, B, i) {
				return true
			}
		} else {
			c := int64(A[i])
			hashA *= BASE
			hashA -= c * rollingBase
			hashA += c
			hashA %= MOD
		}
	}
	return false
}

func main() {
	fmt.Println(rotateString("clrwmpkiru", "wmpkwruclr"))
}
