package kata

import "testing"

func TestOpposite(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 then -1",
			args: args{1},
			want: -1,
		},
		{
			name: "14 then -14",
			args: args{14},
			want: -14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Opposite(tt.args.value); got != tt.want {
				t.Errorf("Opposite() = %v, want %v", got, tt.want)
			}
		})
	}
}
