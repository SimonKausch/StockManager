package main

import "testing"

func Test_parseIntInput(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text    string
		want    int
		wantErr bool
	}{
		{
			name:    "valid integer",
			text:    "450",
			want:    450,
			wantErr: false,
		},
		{
			name:    "invalid integer with character",
			text:    "451h",
			want:    0,
			wantErr: true,
		},
		{
			name:    "float input",
			text:    "1.4533",
			want:    0,
			wantErr: true,
		},
		{
			name:    "negate integer",
			text:    "-30",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parseIntInput(tt.text)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseIntInput() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseIntInput() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("parseIntInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
