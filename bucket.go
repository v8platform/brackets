package raw_parser

import (
	"fmt"
	"github.com/xelaj/go-dry"
	"strings"
)

type BucketItem interface {
	Get(idx int) (BucketItem, bool)
	Len() int
	GetString() string
	Items() ([]BucketItem, bool)
	String() string
}

type Bucket []BucketItem

func (b Bucket) Get(idx int) (interface{}, bool) {

	return b.GetItem(idx)
}

func (b Bucket) GetItem(idx int) (BucketItem, bool) {

	if idx < 0 || b.Len() <= idx {
		return nil, false
	}

	return b[idx], true

}

func (b Bucket) Len() int {
	return len(b)
}

type arrayItem []BucketItem

func (b arrayItem) Get(idx int) (BucketItem, bool) {

	if idx < 0 || b.Len() <= idx {
		return nil, false
	}

	return b[idx], true

}
func (b arrayItem) GetString() string {
	return b.String()
}

func (b arrayItem) String() string {

	var strs []string

	for _, item := range b {
		strs = append(strs, item.String())
	}

	val := strings.Join(strs, ",")

	return fmt.Sprintf("{%s}", val)
}

func (b arrayItem) Len() int {
	return len(b)
}

func (b arrayItem) Items() ([]BucketItem, bool) {
	return []BucketItem{b}, true
}

type stringItem string

func (b stringItem) Get(_ int) (BucketItem, bool) {
	return b, true
}

func (b stringItem) GetString() string {
	return b.String()
}

func (b stringItem) Len() int {
	return 1
}

func (b stringItem) String() string {
	return string(b)
}

func (b stringItem) Items() ([]BucketItem, bool) {
	return []BucketItem{b}, true
}

type BucketsNode struct {
	Text string

	Nodes []BucketsNode
	//Count int

	valueNode bool
}

func (b BucketsNode) GetNode(address ...int) BucketsNode {

	currentNode := b

	for i, _ := range address {
		currentNode = currentNode.Nodes[address[i]]
	}

	return currentNode
}

func (b BucketsNode) String() string {

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

func (b BucketsNode) Int() int {

	return dry.StringToInt(b.Text)
}

func (b BucketsNode) Bool() bool {

	return dry.StringToBool(b.Text)
}

func (b BucketsNode) Float64() float64 {

	return dry.StringToFloat(b.Text)
}

func NewValueNode(value string) BucketsNode {

	return BucketsNode{
		Text:      value,
		valueNode: true,
	}
}

func NewNode() BucketsNode {

	return BucketsNode{
		valueNode: false,
	}
}
