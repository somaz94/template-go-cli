package example

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{"with name", "Go", "Hello, Go!"},
		{"empty name", "", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
