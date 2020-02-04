package main

import "testing"

func TestSumArgs(t *testing.T) {
	// type args struct {
	// 	args []string
	// }
	tests := []struct {
		name string
		args []string
		exp  string
		sum  float64
	}{
		// TODO: Add test cases.
		{
			name: "most postive case",
			args: []string{"1", "3", "5"},
			exp:  "1+3+5",
			sum:  9,
		},
		{
			name: "with invalid number",
			args: []string{"1", "hello", "5"},
			exp:  "1+5",
			sum:  6,
		},
		{
			name: "with invalid number and minus number",
			args: []string{"10", "-2", "10", "-10", "2", "hello", "5", "3"},
			exp:  "10-2+10-10+2+5+3",
			sum:  18,
		},
		{
			name: "with invalid number and start with minus number",
			args: []string{"-10", "-2", "10", "-10", "2", "hello", "5", "3"},
			exp:  "-10-2+10-10+2+5+3",
			sum:  -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp, sum := SumArgs(tt.args)
			if exp != tt.exp {
				t.Errorf("SumArgs() got = %v, want %v", exp, tt.exp)
			}
			if sum != tt.sum {
				t.Errorf("SumArgs() got1 = %v, want %v", sum, tt.sum)
			}
		})
	}
}
