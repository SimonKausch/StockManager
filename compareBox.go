package main

func fitsIn(b BoundingBox, s Stock) bool {
	// Get dimensions of the stock
	sx, sy, sz := float64(s.XLength), float64(s.YLength), float64(s.ZLength)

	// Get dimensions of the bounding box
	b1, b2, b3 := b.BoxX, b.BoxY, b.BoxZ

	// Permutation 1: inner (l1, l2, l3) vs outer (L1, L2, L3)
	if b1 <= sx && b2 <= sy && b3 <= sz {
		return true
	}
	// Permutation 2: inner (l1, l3, l2) vs outer (L1, L2, L3)
	if b1 <= sx && b3 <= sy && b2 <= sz {
		return true
	}
	// Permutation 3: inner (l2, l1, l3) vs outer (L1, L2, L3)
	if b2 <= sx && b1 <= sy && b3 <= sz {
		return true
	}
	// Permutation 4: inner (l2, l3, l1) vs outer (L1, L2, L3)
	if b2 <= sx && b3 <= sy && b1 <= sz {
		return true
	}
	// Permutation 5: inner (l3, l1, l2) vs outer (L1, L2, L3)
	if b3 <= sx && b1 <= sy && b2 <= sz {
		return true
	}
	// Permutation 6: inner (l3, l2, l1) vs outer (L1, L2, L3)
	if b3 <= sx && b2 <= sy && b1 <= sz {
		return true
	}

	return false // return b.BoxX <= float64(s.XLength) &&
	// 	b.BoxY <= float64(s.YLength) &&
	// 	b.BoxZ <= float64(s.ZLength)
}

func findFittingStock(b BoundingBox, s []Stock) []Stock {
	var res []Stock

	for _, i := range s {
		if fitsIn(b, i) {
			res = append(res, i)
		}
	}

	return res
}
