package kata

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[\"arp\", \"live\", \"strong\"] [\"lively\", \"alive\", \"harp\", \"sharp\", \"armstrong\"] then [\"arp\", \"live\", \"strong\"]",
			args: args{[]string{"arp", "live", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}},
			want: []string{"arp", "live", "strong"},
		},
		{
			name: "[\"tarp\", \"mice\", \"bull\"] [\"lively\", \"alive\", \"harp\", \"sharp\", \"armstrong\"] then []",
			args: args{[]string{"tarp", "mice", "bull"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
