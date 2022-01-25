package utils

func Reduce(combine func(m1 Matrix, m2 Matrix) Matrix, m ...Matrix) Matrix {
	var acc = m[0]

	for i := 1; i < len(m); i++ {
		acc = combine(acc, m[i])
	}

	return acc
}
