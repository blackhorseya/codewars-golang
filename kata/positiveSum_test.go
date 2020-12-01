package kata

import "testing"

func TestPositiveSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1, 2, 3, 4, 5] then 15",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "[1, -2, 3, 4, 5] then 13",
			args: args{[]int{1, -2, 3, 4, 5}},
			want: 13,
		},
		{
			name: "[] then 0",
			args: args{[]int{}},
			want: 0,
		},
		{
			name: "[-1, -2, -3, -4, -5] then 0",
			args: args{[]int{-1, -2, -3, -4, -5}},
			want: 0,
		},
		{
			name: "[-1, 2, 3, 4, -5] then 9",
			args: args{[]int{-1, 2, 3, 4, -5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositiveSum(tt.args.numbers); got != tt.want {
				t.Errorf("PositiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
