package kata

import "testing"

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 then Odd",
			args: args{1},
			want: "Odd",
		},
		{
			name: "2 then Even",
			args: args{2},
			want: "Even",
		},
		{
			name: "-1 then Odd",
			args: args{-1},
			want: "Odd",
		},
		{
			name: "-2 then Even",
			args: args{-2},
			want: "Even",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd(tt.args.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
