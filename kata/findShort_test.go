package kata

import "testing"

func TestFindShort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "\"bitcoin take over the world maybe who knows perhaps\" then 3",
			args: args{"bitcoin take over the world maybe who knows perhaps"},
			want: 3,
		},
		{
			name: "\"turns out random test cases are easier than writing out basic ones\" then 3",
			args: args{"turns out random test cases are easier than writing out basic ones"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShort(tt.args.s); got != tt.want {
				t.Errorf("FindShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
