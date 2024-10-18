package string_utils

import "testing"

func TestRevertStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test_1",
			s:    "abc",
			want: "cba",
		},
		{
			name: "Test_2",
			s:    "aaaab",
			want: "baaaa",
		},
		{
			name: "Test_3",
			s:    "a",
			want: "a",
		},
		{
			name: "Test_4",
			s:    "",
			want: "",
		},
		{
			name: "Test_5",
			s:    "АБВ",
			want: "ВБА",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RevertStr(tt.s); got != tt.want {
				t.Errorf("RevertStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
