package kata

import "testing"

func TestDigitalRoot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "16 then 7",
			args: args{16},
			want: 7,
		},
		{
			name: "942 then 6",
			args: args{942},
			want: 6,
		},
		{
			name: "132189 then 6",
			args: args{132189},
			want: 6,
		},
		{
			name: "493193 then 2",
			args: args{493193},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitalRoot(tt.args.n); got != tt.want {
				t.Errorf("DigitalRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
