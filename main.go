package main

import (
	"fmt"
	"github.com/OdaDaisuke/go_utils/node"
	"github.com/OdaDaisuke/go_utils/sliceutil"
	"strconv"
)

func main() {
	fmt.Println(sliceutil.MakeMultiDimension(2, 2))

	n := &node.Node{
		Value: "1",
		Prev:  nil,
	}
	firstNode := n

	for i := 2; i < 10; i++ {
		value := strconv.Itoa(i)
		n = node.Add(n, &node.Node{
			Value: value,
		})
	}

	node.PrintAll(firstNode)
}
