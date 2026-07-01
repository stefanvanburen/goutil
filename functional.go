package goutil

// Blatantly ripped from https://gobyexample.com/collection-functions

import "slices"

// Index returns the first index of the target string `t`, or -1 if no match is found.
//
// Deprecated: use [slices.Index] instead.
//
//go:fix inline
func Index[K comparable](vs []K, t K) int {
	return slices.Index(vs, t)
}

// Include returns `true` if the target string t is in the slice.
//
// Deprecated: use [slices.Contains] instead.
//
//go:fix inline
func Include[K comparable](vs []K, t K) bool {
	return slices.Contains(vs, t)
}

// Any returns `true` if one of the Ks in the slice satisfies the predicate `f`.
//
// Deprecated: use [slices.ContainsFunc] instead.
//
//go:fix inline
func Any[K comparable](vs []K, f func(K) bool) bool {
	return slices.ContainsFunc(vs, f)
}

// All returns `true` if all of the Ks in the slice satisfy the predicate `f`.
func All[K comparable](vs []K, f func(K) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all Ks in the slice that satisfy
// the predicate `f`.
func Filter[K comparable](vs []K, f func(K) bool) []K {
	var vsf []K
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function `f`
// to each K in the original slice.
func Map[K comparable](vs []K, f func(K) K) []K {
	vsm := make([]K, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
