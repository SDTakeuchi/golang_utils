package sliceutils

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

func Push[T comparable](s []T, elem T) []T {
	return append(s, elem)
}

func Pop[T comparable](s []T) []T {
	return s[:len(s)-1]
}

func Unshift[T comparable](s []T, elem T) []T {
	return append([]T{elem}, s...)
}

func Shift[T comparable](s []T) []T {
	return s[1:]
}

func Insert[T comparable](s []T, v T, index int) []T {
	var tempVal T
	s = append(s, tempVal)
	s = append(s[:index+1], s[index:len(s)-1]...)
	s[index] = v
	return s
}

func Remove[T comparable](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}