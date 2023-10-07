package search_in_a_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Find the node in the BST that the node's value equals val and return the subtree rooted with that node",
			args: args{
				&TreeNode{4,
					&TreeNode{2,
						&TreeNode{1, nil, nil},
						&TreeNode{3, nil, nil}},
					&TreeNode{7, nil, nil}},
				2},
			want: &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		},
		{
			name: "If such a node does not exist, return nil",
			args: args{
				&TreeNode{4,
					&TreeNode{2,
						&TreeNode{1, nil, nil},
						&TreeNode{3, nil, nil}},
					&TreeNode{7, nil, nil}},
				10},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
