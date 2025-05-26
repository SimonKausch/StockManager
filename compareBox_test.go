package main

import "testing"

// Test suite
func TestFitsIn(t *testing.T) {
	testCases := []struct {
		name     string
		box      BoundingBox
		stock    Stock
		expected bool
	}{
		{
			name:     "Box fits perfectly",
			box:      BoundingBox{BoxX: 10.0, BoxY: 20.0, BoxZ: 30.0},
			stock:    Stock{ID: 1, XLength: 10, YLength: 20, ZLength: 30},
			expected: true,
		},
		{
			name:     "Box is smaller and fits",
			box:      BoundingBox{BoxX: 5.0, BoxY: 15.0, BoxZ: 25.0},
			stock:    Stock{ID: 2, XLength: 10, YLength: 20, ZLength: 30},
			expected: true,
		},
		{
			name:     "Box is smaller but oriented wrong",
			box:      BoundingBox{BoxX: 15.0, BoxY: 9.0, BoxZ: 25.0},
			stock:    Stock{ID: 2, XLength: 10, YLength: 20, ZLength: 30},
			expected: true,
		},
		{
			name:     "Box is larger in X dimension",
			box:      BoundingBox{BoxX: 11.0, BoxY: 20.0, BoxZ: 30.0},
			stock:    Stock{ID: 3, XLength: 10, YLength: 20, ZLength: 30},
			expected: false,
		},
		{
			name:     "Box is larger in Y dimension",
			box:      BoundingBox{BoxX: 10.0, BoxY: 21.0, BoxZ: 30.0},
			stock:    Stock{ID: 4, XLength: 10, YLength: 20, ZLength: 30},
			expected: false,
		},
		{
			name:     "Box is larger in Z dimension",
			box:      BoundingBox{BoxX: 10.0, BoxY: 20.0, BoxZ: 31.0},
			stock:    Stock{ID: 5, XLength: 10, YLength: 20, ZLength: 30},
			expected: false,
		},
		{
			name:     "Box has zero dimensions, stock has dimensions",
			box:      BoundingBox{BoxX: 0.0, BoxY: 0.0, BoxZ: 0.0},
			stock:    Stock{ID: 8, XLength: 10, YLength: 20, ZLength: 30},
			expected: true,
		},
		{
			name:     "Box has dimensions, stock has zero dimensions for X",
			box:      BoundingBox{BoxX: 1.0, BoxY: 1.0, BoxZ: 1.0},
			stock:    Stock{ID: 9, XLength: 0, YLength: 1, ZLength: 1},
			expected: false,
		},
		{
			name:     "Box has dimensions, stock has zero dimensions for all",
			box:      BoundingBox{BoxX: 1.0, BoxY: 1.0, BoxZ: 1.0},
			stock:    Stock{ID: 10, XLength: 0, YLength: 0, ZLength: 0},
			expected: false,
		},
		{
			name:     "Box and stock have zero dimensions",
			box:      BoundingBox{BoxX: 0.0, BoxY: 0.0, BoxZ: 0.0},
			stock:    Stock{ID: 11, XLength: 0, YLength: 0, ZLength: 0},
			expected: true,
		},
		{
			name:     "Box dimension slightly larger (float precision test for X)",
			box:      BoundingBox{BoxX: 10.000000000000001, BoxY: 20.0, BoxZ: 30.0}, // Smallest value > 10
			stock:    Stock{ID: 12, XLength: 10, YLength: 20, ZLength: 30},
			expected: false,
		},
		{
			name:     "Box dimension equal (float precision test for X)",
			box:      BoundingBox{BoxX: 10.0, BoxY: 20.0, BoxZ: 30.0},
			stock:    Stock{ID: 13, XLength: 10, YLength: 20, ZLength: 30},
			expected: true,
		},
		{
			name:     "Stock has negative dimensions, box is positive (should not fit if interpreted strictly)",
			box:      BoundingBox{BoxX: 5.0, BoxY: 5.0, BoxZ: 5.0},
			stock:    Stock{ID: 14, XLength: -10, YLength: 20, ZLength: 30}, // Assuming lengths are non-negative
			expected: false,                                                 // b.BoxX (5.0) <= float64(-10) is false
		},
		{
			name: "Box has negative dimensions, stock is positive (can fit if interpreted as position)",
			box:  BoundingBox{BoxX: -5.0, BoxY: -5.0, BoxZ: -5.0}, // If dimensions are always positive, this case might be invalid input.
			// However, the function will evaluate it based on comparison.
			stock:    Stock{ID: 15, XLength: 10, YLength: 20, ZLength: 30},
			expected: true, // -5.0 <= 10.0 is true
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fitsIn(tc.box, tc.stock)
			if result != tc.expected {
				t.Errorf("fitsIn(%+v, %+v) = %t; want %t", tc.box, tc.stock, result, tc.expected)
			}
		})
	}
}
