package helperFunctions

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Keys return keys of any map
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type Set[T comparable] map[T]struct{}

// Make returns a set of some element type.
func Make[T comparable](x ...int) Set[T] {
	if len(x) > 0 {
		n := x[0]
		return make(Set[T], n)
	}
	return make(Set[T])
}

// Add adds v to the set s.
// If v is already in s this has no effect.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Delete removes v from the set s.
// If v is not in s this has no effect.
func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Contains reports whether v is in s.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Len reports the number of elements in s.
func (s Set[T]) Len() int {
	return len(s)
}

// Iterate invokes f on each element of s.
// It's OK for f to call the Delete method.
func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}
