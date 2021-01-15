package raw_parser

import (
	"reflect"
	"testing"
)

func Test_bucketsNode_GetNode(t *testing.T) {
	type fields struct {
		Text      string
		Nodes     BracketNodes
		valueNode bool
	}
	type args struct {
		address []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BracketsNode
		wantErr bool
	}{
		{
			"simple",
			fields{
				Nodes: BracketNodes{
					newValueNode("0"),
					bracketsNode{
						Nodes: BracketNodes{
							newValueNode("1"),
						},
					},
				},
			},
			args{
				address: []int{1, 0},
			},
			newValueNode("1"),
			false,
		},
		{
			"error",
			fields{
				Nodes: BracketNodes{
					newValueNode("0"),
					bracketsNode{
						Nodes: BracketNodes{
							newValueNode("1"),
						},
					},
				},
			},
			args{
				address: []int{1, 1},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bracketsNode{
				Text:      tt.fields.Text,
				Nodes:     tt.fields.Nodes,
				valueNode: tt.fields.valueNode,
			}
			got, err := b.GetNode(tt.args.address...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
