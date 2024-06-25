package main

import (
	"reflect"
	"testing"
)

func Test_charCount(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			if got := charCount(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(i)
	}
}

func BenchmarkFib1(b *testing.B) { benchFib(1, b)}
func BenchmarkFib10(b *testing.B) { benchFib(10, b)}
func BenchmarkFib20(b *testing.B) { benchFib(20, b)}