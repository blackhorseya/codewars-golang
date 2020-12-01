package kata

import "testing"

func TestHighAndLow(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[8 3 -5 42 -1 0 0 -9 4 7 4 -4] then [42 -9]",
			args: args{"8 3 -5 42 -1 0 0 -9 4 7 4 -4"},
			want: "42 -9",
		},
		{
			name: "[1 2 -3 4 5] then [5 -3]",
			args: args{"1 2 -3 4 5"},
			want: "5 -3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighAndLow(tt.args.in); got != tt.want {
				t.Errorf("HighAndLow() = %v, want %v", got, tt.want)
			}
		})
	}
}
