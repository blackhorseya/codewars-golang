package kata

import "testing"

func TestSpinWords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"CodeWars\" then \"sraWedoC\"",
			args: args{"CodeWars"},
			want: "sraWedoC",
		},
		{
			name: "\"to\" then \"to\"",
			args: args{"to"},
			want: "to",
		},
		{
			name: "\"Hey fellow warriors\" then \"Hey wollef sroirraw\"",
			args: args{"Hey fellow warriors"},
			want: "Hey wollef sroirraw",
		},
		{
			name: "\"Burgers are my favorite fruit\" then \"sregruB are my etirovaf tiurf\"",
			args: args{"Burgers are my favorite fruit"},
			want: "sregruB are my etirovaf tiurf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpinWords(tt.args.str); got != tt.want {
				t.Errorf("SpinWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
