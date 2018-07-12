package main

import (
	"fmt"

	"github.com/johnnyluo/basic-algorithm/list"
)

func main() {
	ll := list.NewSingleLinkList()
	for i := 0; i < 100; i++ {
		if err := ll.AddNode(i); nil != err {
			panic(err)
		}
	}
	if err := ll.Traverse(func(node *list.SingleLinkNode) {
		fmt.Printf("%d ", node.Value)
	}); nil != err {
		panic(err)
	}
	if err := ll.Reverse(); nil != err {
		panic(err)
	}
	fmt.Println("===================+++++++")
	if err := ll.Traverse(func(node *list.SingleLinkNode) {
		fmt.Printf("%d ", node.Value)
	}); nil != err {
		panic(err)
	}
}
