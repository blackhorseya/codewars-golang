package kata

import "testing"

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"hello case\" then \"HelloCase\"",
			args: args{"hello case"},
			want: "HelloCase",
		},
		{
			name: "\"camel case method\" then \"CamelCaseMethod\"",
			args: args{"camel case method"},
			want: "CamelCaseMethod",
		},
		{
			name: "\"\" then \"\"",
			args: args{""},
			want: "",
		},
		{
			name: "\" camel case word\" then \"CamelCaseWord\"",
			args: args{" camel case word"},
			want: "CamelCaseWord",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
