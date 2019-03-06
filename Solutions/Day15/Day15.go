package main

import(
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func(list *LinkedList) Insert(item int) {
	newElement := &Node{data:item}
	if list.head == nil{
		list.head = newElement
	} else {
		start := list.head
		for list.head.next!=nil{
			list.head=list.head.next
		}
		list.head.next = newElement
		list.head=start
	}
}

func(list *LinkedList) Display(){
	start := list.head
	if list.head ==nil{
		return
	}
	for list.head!= nil{
			fmt.Printf("%d ",list.head.data)
			list.head=list.head.next
		}
	list.head = start

}


func NewList() *LinkedList {
	return &LinkedList{head:nil}
}

func main() {
	L := NewList()
	var t int
	var d int
	fmt.Scanf("%d\n",&t)
	for i:=0;i<t;i++{
		fmt.Scanf("%d\n",&d)
		L.Insert(d)
	}
	L.Display()
}
