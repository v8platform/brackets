package brackets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Node интерфес для работы с нодой
// 	Нода представляет собой структуру типа:
//
//	{
//		0,
//      {0,0},
//		"Text",
//		{"text", 0}
//	}
//
type Node interface {

	// GetNode получение вложенных нов по их адресу в массиве норд
	GetNode(address ...int) Node
	GetNodeE(address ...int) (Node, error)

	// Get получение значение ноды по адресу
	Get(address ...int) string

	// String получение строкового значение ноды
	//	Для нод без вложенных нод выводиться их значение
	//	Для всех остальных выводиться строковое представление ноды
	String() string

	// Int64 получение числового значение ноды
	Int(address ...int) int

	// Int64 получение числового значение ноды
	Int64(address ...int) int64
	// Bool получение булевного значение ноды
	Bool(address ...int) bool
	// Int получение числового с запятой значение ноды
	Float64(address ...int) float64

	value() bool
}

var ErrNodeAddress = errors.New("address node is broken")
var emptyNode = bracketsNode{}

func EmptyNode(n Node) bool {

	bNode, ok := n.(bracketsNode)

	if !ok {
		return false
	}

	if len(bNode.Text) > 0 || len(bNode.Nodes) > 0 {
		return true
	}

	return false
}

// Nodes Массив Node с функцией String()
type Nodes []Node

func (b Nodes) String() string {

	var strs []string

	for _, item := range b {
		strs = append(strs, item.String())
	}

	val := strings.Join(strs, ",")

	return fmt.Sprintf("{%s}", val)
}

type bracketsNode struct {
	Text  string
	Nodes Nodes
	size      int
	valueNode bool
}

func (b bracketsNode) value() bool {
	return b.valueNode
}

func (b bracketsNode) Get(address ...int) string {

	if len(address) == 0 && b.value() {
		return b.Text
	}

	node, _ := b.getNode(address)
	if node.value() {
		return node.Text
	}

	return ""

}

func (b bracketsNode) getNode(address []int) (bracketsNode, error) {

	currentNode := b

	for i, _ := range address {

		if len(currentNode.Nodes) <= address[i] {
			return emptyNode, ErrNodeAddress
		}

		currentNode = currentNode.Nodes[address[i]].(bracketsNode)
	}

	return currentNode, nil

}

func (b bracketsNode) GetNode(address ...int) Node {

	node, _ := b.getNode(address)
	return node
}

func (b bracketsNode) GetNodeE(address ...int) (Node, error) {

	return b.getNode(address)
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

func (b bracketsNode) Int(address ...int) int {

	return int(b.Int64(address...))
}

func (b bracketsNode) Int64(address ...int) int64 {

	n, err := b.getNode(address)
	if err != nil ||
		!n.value() ||
		len(n.Text) == 0 {
		return 0
	}

	i, _ := strconv.ParseInt(n.Text, 10, 64)
	return i
}

func (b bracketsNode) Bool(address ...int) bool {

	n, err := b.getNode(address)
	if err != nil ||
		!n.value() ||
		len(n.Text) == 0 {
		return false
	}
	val, _ := strconv.ParseBool(n.Text)
	return val
}

func (b bracketsNode) Float64(address ...int) float64 {

	n, err := b.getNode(address)
	if err != nil ||
		!n.value() ||
		len(n.Text) == 0 {
		return 0
	}

	f, _ := strconv.ParseFloat(n.Text, 64)
	return f
}

func newValueNode(value string) bracketsNode {
	addQuotesIfNeed(&value)
	return bracketsNode{
		Text:      value,
		valueNode: true,
		size:      len([]byte(value)),
	}
}

func addQuotesIfNeed(value *string) {

	if strings.ContainsAny(*value, `'"`) {
		*value = `"` + *value + `"`
	}
}
