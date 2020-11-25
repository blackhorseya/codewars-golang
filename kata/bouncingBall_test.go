package kata

import "testing"

func TestBouncingBall(t *testing.T) {
	type args struct {
		h      float64
		bounce float64
		window float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 0.6 10 then -1",
			args: args{0, 0.6, 10},
			want: -1,
		},
		{
			name: "10 1.2 10 then -1",
			args: args{10, 1.2, 10},
			want: -1,
		},
		{
			name: "10 0.6 10 then -1",
			args: args{10, 0.6, 10},
			want: -1,
		},
		{
			name: "3 0.66 1.5 then 3",
			args: args{3, 0.66, 1.5},
			want: 3,
		},
		{
			name: "40 0.4 10 then 3",
			args: args{40, 0.4, 10},
			want: 3,
		},
		{
			name: "4 0.25 1 then 1",
			args: args{4, 0.25, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BouncingBall(tt.args.h, tt.args.bounce, tt.args.window); got != tt.want {
				t.Errorf("BouncingBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
