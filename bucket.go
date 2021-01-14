package raw_parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BucketsNode interface {
	GetNode(address ...int) (BucketsNode, error)

	String() string
	Int() int
	Bool() bool
	Float64() float64
}

var ErrNodeAddress = errors.New("address node is broken")

type BucketNodes []BucketsNode

func (b BucketNodes) String() string {

	var strs []string

	for _, item := range b {
		strs = append(strs, item.String())
	}

	val := strings.Join(strs, ",")

	return fmt.Sprintf("{%s}", val)
}

type bucketsNode struct {
	Text string

	Nodes BucketNodes
	//Count int

	valueNode bool
}

func (b bucketsNode) GetNode(address ...int) (BucketsNode, error) {

	currentNode := b

	for i, _ := range address {

		if len(currentNode.Nodes) <= address[i] {
			return nil, ErrNodeAddress
		}

		currentNode = currentNode.Nodes[address[i]].(bucketsNode)
	}

	return currentNode, nil
}

func (b bucketsNode) String() string {

	if b.valueNode {
		return b.Text
	}

	var strs []string

	for _, item := range b.Nodes {
		strs = append(strs, item.String())
	}

	val := strings.Join(strs, ",")

	return fmt.Sprintf("{%s}", val)
}

func (b bucketsNode) Int() int {
	i, _ := strconv.ParseInt(b.Text, 10, 64)
	return int(i)
}

func (b bucketsNode) Bool() bool {

	val, _ := strconv.ParseBool(b.Text)
	return val
}

func (b bucketsNode) Float64() float64 {

	f, _ := strconv.ParseFloat(b.Text, 64)
	return f
}

func NewValueNode(value string) bucketsNode {

	return bucketsNode{
		Text:      value,
		valueNode: true,
	}
}

func NewNode() bucketsNode {

	return bucketsNode{}
}
