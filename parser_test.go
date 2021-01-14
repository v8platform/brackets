package raw_parser

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestBucketsParser_nextNodeText(t *testing.T) {

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			"simple",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := BucketsParser{
				rd: bufio.NewReader(strings.NewReader(tt.text)),
			}
			if got := p.nextNodeText(); got != tt.want {
				t.Errorf("nextNodeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketsParser_NextNode(t *testing.T) {
	tests := []struct {
		name string
		text string
		want *BucketsNode
	}{
		{
			"simple",
			"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n},\n{20200412134356,N,\n{0,0},1,1,2,2,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,2,0,\n{0}\n},",
			&BucketsNode{
				Nodes: []BucketsNode{
					NewValueNode("20200412134348"),
					NewValueNode("N"),
					{
						Nodes: []BucketsNode{
							NewValueNode("0"),
							NewValueNode("0"),
						},
					},
					NewValueNode("1"),
				},
			}, //"{20200412134348,N,\n{0,0},1,1,1,1,1,I,\"\",0,\n{\"U\"},\"\",1,1,0,1,0,\n{0}\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := BucketsParser{
				rd: bufio.NewReader(strings.NewReader(tt.text)),
			}
			if got := p.NextNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
