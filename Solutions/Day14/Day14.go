package main

import (
	"fmt"
	"sort"
)

type Difference struct {
	elements          []int
	maximumDiffernece int
}

func (d *Difference) ComputeDifference() int {
	sort.Ints(d.elements)
	d.maximumDiffernece = d.elements[len(d.elements)-1] - d.elements[0]
	return d.maximumDiffernece
}

func main() {
	var s int
	var d int
	fmt.Scanf("%d\n",&s)
	arr := make([]int,s)
	for i:=0;i<s;i++{
		fmt.Scanf("%d\n",&d)
		arr[i]=d
	}
	diff := Difference{arr,0}
	fmt.Println(diff.ComputeDifference())

}
