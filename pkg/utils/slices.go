package utils

func Delete[S ~[]E, E any](s S, i, j int) S {
	_ = s[i:j:len(s)] // bounds check

	if i == j {
		return s
	}

	oldlen := len(s)
	s = append(s[:i], s[j:]...)
	clearSlice(s[len(s):oldlen]) // zero/nil out the obsolete elements, for GC
	return s
}

func clearSlice[S ~[]E, E any](s S) {
	var zero E
	for i := range s {
		s[i] = zero
	}
}
