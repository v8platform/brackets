package raw_parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BracketsNode interface {
	GetNode(address ...int) (BracketsNode, error)

	String() string
	Int() int
	Bool() bool
	Float64() float64
}

var ErrNodeAddress = errors.New("address node is broken")

type BracketNodes []BracketsNode

func (b BracketNodes) String() string {

	var strs []string

	for _, item := range b {
		strs = append(strs, item.String())
	}

	val := strings.Join(strs, ",")

	return fmt.Sprintf("{%s}", val)
}

type bracketsNode struct {
	Text string

	Nodes BracketNodes
	//Count int

	valueNode bool
}

func (b bracketsNode) GetNode(address ...int) (BracketsNode, error) {

	currentNode := b

	for i, _ := range address {

		if len(currentNode.Nodes) <= address[i] {
			return nil, ErrNodeAddress
		}

		currentNode = currentNode.Nodes[address[i]].(bracketsNode)
	}

	return currentNode, nil
}

func (b bracketsNode) String() string {

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

func (b bracketsNode) Int() int {
	i, _ := strconv.ParseInt(b.Text, 10, 64)
	return int(i)
}

func (b bracketsNode) Bool() bool {

	val, _ := strconv.ParseBool(b.Text)
	return val
}

func (b bracketsNode) Float64() float64 {

	f, _ := strconv.ParseFloat(b.Text, 64)
	return f
}

func newValueNode(value string) bracketsNode {

	return bracketsNode{
		Text:      value,
		valueNode: true,
	}
}
