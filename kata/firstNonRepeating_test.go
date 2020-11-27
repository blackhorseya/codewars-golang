package kata

import "testing"

func Test_firstNonRepeating(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "\"a\" then \"a\"",
			args: args{"a"},
			want: "a",
		},
		{
			name: "\"stress\" then \"t\"",
			args: args{"stress"},
			want: "t",
		},
		{
			name: "\"moonmen\" then \"e\"",
			args: args{"moonmen"},
			want: "e",
		},
		{
			name: "\"~><#~><\" then \"#\"",
			args: args{"~><#~><"},
			want: "#",
		},
		{
			name: "\"sTreSS\" then \"T\"",
			args: args{"sTreSS"},
			want: "T",
		},
		{
			name: "\"Go hang a salami, I'm a lasagna hog!\" then \",\"",
			args: args{"Go hang a salami, I'm a lasagna hog!"},
			want: ",",
		},
		{
			name: "\"\" then \"\"",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonRepeating(tt.args.str); got != tt.want {
				t.Errorf("firstNonRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
