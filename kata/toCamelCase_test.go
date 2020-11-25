package kata

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "the-stealth-warrior then theStealthWarrior",
			args: args{"the-stealth-warrior"},
			want: "theStealthWarrior",
		},
		{
			name: "The_Stealth_Warrior then TheStealthWarrior",
			args: args{"The_Stealth_Warrior"},
			want: "TheStealthWarrior",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.s); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
