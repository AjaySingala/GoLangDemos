package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

}
