package kata

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "\"a\" then 0",
			args: args{"a"},
			want: 0,
		},
		{
			name: "\"bcd\" then 9",
			args: args{"bcd"},
			want: 9,
		},
		{
			name: "\"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes\" then 143",
			args: args{"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"},
			want: 143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.str); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
