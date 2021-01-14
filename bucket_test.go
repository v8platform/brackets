package raw_parser

import (
	"reflect"
	"testing"
)

func Test_bucketsNode_GetNode(t *testing.T) {
	type fields struct {
		Text      string
		Nodes     BucketNodes
		valueNode bool
	}
	type args struct {
		address []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BucketsNode
		wantErr bool
	}{
		{
			"simple",
			fields{
				Nodes: BucketNodes{
					NewValueNode("0"),
					bucketsNode{
						Nodes: BucketNodes{
							NewValueNode("1"),
						},
					},
				},
			},
			args{
				address: []int{1, 0},
			},
			NewValueNode("1"),
			false,
		},
		{
			"error",
			fields{
				Nodes: BucketNodes{
					NewValueNode("0"),
					bucketsNode{
						Nodes: BucketNodes{
							NewValueNode("1"),
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
			b := bucketsNode{
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
