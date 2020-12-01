package kata

import "testing"

func TestGetCount(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name:      "\"abracadabra\" then 5",
			args:      args{"abracadabra"},
			wantCount: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := GetCount(tt.args.str); gotCount != tt.wantCount {
				t.Errorf("GetCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
