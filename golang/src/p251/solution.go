package main

import (
	"fmt"
)

type Vector2D struct {
	v [][]int
	i, j int
}


func Constructor(v [][]int) Vector2D {
	return Vector2D{
		v: v,
		i: 0,
		j: 0,
	}
}

func (this *Vector2D) adjust() {
	for this.i < len(this.v) && this.j == len(this.v[this.i]) {
		this.i++
		this.j = 0
	}
}

func (this *Vector2D) Next() (r int) {
	r = this.v[this.i][this.j]
	this.j++
	this.adjust()
	return
}


func (this *Vector2D) HasNext() bool {
	if len(this.v) == 0 {
		return false
	}

	this.adjust()
	return this.i < len(this.v) && this.j != len(this.v[this.i])
}

func main() {
	v := Constructor([][]int{
		{1, 2},
		{3},
		{4},
	})
	for v.HasNext() {
		fmt.Println(v.Next())
		fmt.Println(v.Next())
		fmt.Println(v.Next())
		fmt.Println(v.HasNext())
		fmt.Println(v.HasNext())
		fmt.Println(v.Next())
		fmt.Println(v.HasNext())
	}
}
