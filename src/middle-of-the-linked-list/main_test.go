package middle_of_the_linked_list

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "return the middle node of the linked list",
			args: args{&ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4, Next: &ListNode{
								Val: 5, Next: nil,
							},
						},
					},
				},
			}},
			want: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil,
					},
				},
			},
		},
		{
			name: "If there are two middle nodes, return the second middle node.",
			args: args{&ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4, Next: nil,
						},
					},
				},
			}},
			want: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
