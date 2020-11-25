package kata

import (
	"reflect"
	"testing"
)

func TestRace(t *testing.T) {
	type args struct {
		v1 int
		v2 int
		g  int
	}
	tests := []struct {
		name string
		args args
		want [3]int
	}{
		{
			name: "820 81 550 then [-1, -1, -1]",
			args: args{820, 81, 550},
			want: [3]int{-1, -1, -1},
		},
		{
			name: "720 850 70 then [0, 32, 18]",
			args: args{720, 850, 70},
			want: [3]int{0, 32, 18},
		},
		{
			name: "80 91 37 then [3, 21, 49]",
			args: args{80, 91, 37},
			want: [3]int{3, 21, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Race(tt.args.v1, tt.args.v2, tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Race() = %v, want %v", got, tt.want)
			}
		})
	}
}
