package foreach

import (
	"reflect"
	"testing"
)

func TestForeachGraph(t *testing.T) {
	type args struct {
		graphNodes []Node
		graphEdges []Edge
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "流水线",
			args: args{
				graphNodes: []Node{
					{"a"},
					{"b"},
					{"c"},
				},
				graphEdges: []Edge{
					{"a", "b"},
					{"b", "c"},
				},
			},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name: "dag(出度大于1)",
			args: args{
				graphNodes: []Node{
					{"a"},
					{"b"},
					{"c"},
					{"d"},
				},
				graphEdges: []Edge{
					{"a", "b"},
					{"b", "c"},
					{"b", "d"},
				},
			},
			want:    []string{"a", "b", "c,d"}, // 也可能是 d,c 这一个测试样例不一定 100% 通过
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ForeachGraph(tt.args.graphNodes, tt.args.graphEdges)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForeachGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForeachGraph() got = %v, want %v", got, tt.want)
			}
		})
	}
}
