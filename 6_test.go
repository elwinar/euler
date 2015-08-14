package main

import (
	"testing"
)

func TestBrute(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{0, 0},
		{1, 0},
		{2, 4},
		{3, 22},
	}

	for n, c := range cases {
		output := Brute(c.input)
		if output != c.output {
			t.Error("case", n, "fail. Expected:", c.output, ". Got:", output, ".")
		}
	}
}

func BenchmarkBrute100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Brute(100)
	}
}

func BenchmarkBrute1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Brute(1000)
	}
}

func BenchmarkBrute10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Brute(10000)
	}
}

func TestElegant(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{0, 0},
		{1, 0},
		{2, 4},
		{3, 22},
	}

	for n, c := range cases {
		output := Elegant(c.input)
		if output != c.output {
			t.Error("case", n, "fail. Expected:", c.output, ". Got:", output, ".")
		}
	}
}

func BenchmarkElegant100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Elegant(100)
	}
}

func BenchmarkElegant1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Elegant(1000)
	}
}

func BenchmarkElelegant10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Elegant(10000)
	}
}
