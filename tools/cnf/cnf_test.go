package cnf

import "testing"

func Test_readCnf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadCnf()
		})
	}
}
