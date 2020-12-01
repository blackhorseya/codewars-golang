package kata

import "testing"

func TestDNAStrand(t *testing.T) {
	type args struct {
		dna string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"ATTGC\" then \"TAACG\"",
			args: args{"ATTGC"},
			want: "TAACG",
		},
		{
			name: "\"AAAA\" then \"TTTT\"",
			args: args{"AAAA"},
			want: "TTTT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAStrand(tt.args.dna); got != tt.want {
				t.Errorf("DNAStrand() = %v, want %v", got, tt.want)
			}
		})
	}
}
