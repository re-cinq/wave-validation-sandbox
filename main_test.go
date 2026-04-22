package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"zero", 0, 0, 0},
		{"positive", 2, 3, 5},
		{"negative", -4, -6, -10},
		{"mixed", -7, 10, 3},
		{"boundary_max", 1<<31 - 1, 0, 1<<31 - 1},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestGreet(t *testing.T) {
	got := Greet("world")
	want := "hello, world"
	if got != want {
		t.Errorf("Greet(%q) = %q; want %q", "world", got, want)
	}
}
