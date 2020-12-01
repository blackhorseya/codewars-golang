package kata

import "testing"

func TestAccum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abcd then A-Bb-Ccc-Dddd",
			args: args{"abcd"},
			want: "A-Bb-Ccc-Dddd",
		},
		{
			name: "ZpglnRxqenU then Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu",
			args: args{"ZpglnRxqenU"},
			want: "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Accum(tt.args.s); got != tt.want {
				t.Errorf("Accum() = %v, want %v", got, tt.want)
			}
		})
	}
}
