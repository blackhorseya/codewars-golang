package kata

import "testing"

func TestRowSumOddNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 then 1",
			args: args{1},
			want: 1,
		},
		{
			name: "3 then 27",
			args: args{3},
			want: 27,
		},
		{
			name: "13 then 2197",
			args: args{13},
			want: 2197,
		},
		{
			name: "19 then 6859",
			args: args{19},
			want: 6859,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RowSumOddNumbers(tt.args.n); got != tt.want {
				t.Errorf("RowSumOddNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
