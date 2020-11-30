package kata

import "testing"

func TestMultiple3And5(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-10 then 0",
			args: args{-10},
			want: 0,
		},
		{
			name: "10 then 23",
			args: args{10},
			want: 23,
		},
		{
			name: "20 then 78",
			args: args{20},
			want: 78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple3And5(tt.args.number); got != tt.want {
				t.Errorf("Multiple3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
