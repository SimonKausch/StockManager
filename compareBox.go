package main

func fitsIn(b BoundingBox, s Stock) bool {
	// TODO: Check all rotations and round down BoundingBox lengths

	return b.BoxX <= float64(s.XLength) &&
		b.BoxY <= float64(s.YLength) &&
		b.BoxZ <= float64(s.ZLength)
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
