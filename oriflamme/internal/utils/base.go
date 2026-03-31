package utils

import "iter"

func Iterate[E any](e []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := range e {
			if !yield(i, e[i]) {
				return
			}
		}
	}
}
