package math

/*
	Every function in the packages we've seen start with a capital letter.
	In Go if something starts with a capital letter that means other packages
	(and programs) are able to see it. If we had named the function average instead of
	Average our main program would not have been able to see it.
*/
func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}
