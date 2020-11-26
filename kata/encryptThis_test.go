package kata

import "testing"

func TestEncryptThis(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"\" then \"\"",
			args: args{""},
			want: "",
		},
		{
			name: "\"A wise old owl lived in an oak\" then \"65 119esi 111dl 111lw 108dvei 105n 97n 111ka\"",
			args: args{"A wise old owl lived in an oak"},
			want: "65 119esi 111dl 111lw 108dvei 105n 97n 111ka",
		},
		{
			name: "\"The more he saw the less he spoke\" then \"84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp\"",
			args: args{"The more he saw the less he spoke"},
			want: "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptThis(tt.args.text); got != tt.want {
				t.Errorf("EncryptThis() = %v, want %v", got, tt.want)
			}
		})
	}
}
