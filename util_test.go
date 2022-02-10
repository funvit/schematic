package schemagen

import "testing"

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			"User",
			"user",
		},
		{
			"TwoWords",
			"two_words",
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg+"|"+tt.want, func(t *testing.T) {
			if got := CamelToSnake(tt.arg); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			"user",
			"user",
		},
		{
			"two_words",
			"twoWords",
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg+"|"+tt.want, func(t *testing.T) {
			if got := SnakeToCamel(tt.arg); got != tt.want {
				t.Errorf("SnakeToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
