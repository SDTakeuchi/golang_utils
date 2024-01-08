package slices

func IndexOf[T comparable](s []T, elem T) int {
	/*
	returns index of given elem in s
	if elem was not found in s slice, returns -1
	*/
	for i, e := range s {
		if e == elem {
			return i
		}
	}
	return -1
}

func Prepend[T comparable](s []T, elem T) []T {
	return append([]T{elem}, s...)
}

func Pop[T comparable](s []T) []T {
	return s[:len(s)-1]
}

func Shift[T comparable](s []T) []T {
	return s[1:]
}

func Insert[T comparable](s []T, v T, i int) []T {
	var tempVal T
	s = append(s, tempVal)
	s = append(s[:i+1], s[i:len(s)-1]...)
	s[i] = v
	return s
}

func Remove[T comparable](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
