package kata

import "testing"

func TestCreatePhoneNumber(t *testing.T) {
	type args struct {
		numbers [10]uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[1,2,3,4,5,6,7,8,9,0] then \"(123) 456-7890\"",
			args: args{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
			want: "(123) 456-7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatePhoneNumber(tt.args.numbers); got != tt.want {
				t.Errorf("CreatePhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
