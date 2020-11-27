package kata

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		parens string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "\"\" then true",
			args: args{""},
			want: true,
		},
		{
			name: "\"(\" then false",
			args: args{"("},
			want: false,
		},
		{
			name: "\")\" then false",
			args: args{")"},
			want: false,
		},
		{
			name: "\"()\" then true",
			args: args{"()"},
			want: true,
		},
		{
			name: "\")(()))\" then false",
			args: args{")(()))"},
			want: false,
		},
		{
			name: "\"(())((()())())\" then true",
			args: args{"(())((()())())"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.parens); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
