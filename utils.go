package utils

func IndexOf[T comparable](elems []T, elem T) int {
	/*
	returns index of given elem in elems
	if elem is not found in elems slice, returns -1
	*/
	for i, e := range elems {
		if e == elem { return i }
	}
	return -1
}
