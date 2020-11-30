package kata

import "testing"

func TestFindOdd(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "\"20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5\" then 5",
			args: args{[]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}},
			want: 5,
		},
		{
			name: "\"1,1,2,-2,5,2,4,4,-1,-2,5\" then -1",
			args: args{[]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOdd(tt.args.seq); got != tt.want {
				t.Errorf("FindOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
