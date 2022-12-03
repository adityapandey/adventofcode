package util

type Set[T comparable] map[T]struct{}

func SetOf[T comparable](a []T) Set[T] {
	s := make(Set[T])
	for _, c := range a {
		s[c] = struct{}{}
	}
	return s
}

func (s Set[T]) Values() []T {
	v := make([]T, 0, len(s))
	for k := range s {
		v = append(v, k)
	}
	return v
}

func (s Set[T]) Intersect(t Set[T]) Set[T] {
	intersection := make(Set[T])
	for v := range t {
		if _, ok := s[v]; ok {
			intersection[v] = struct{}{}
		}
	}
	return intersection
}

func (s Set[T]) Union(t Set[T]) Set[T] {
	union := make(Set[T])
	for v := range s {
		union[v] = struct{}{}
	}
	for v := range t {
		union[v] = struct{}{}
	}
	return union
}
