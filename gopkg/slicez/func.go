package slicez

// Filter 基于自定义函数转化切片
func Filter[T any](src []T, f func(T) bool) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

// Map 基于自定义函数映射切片
func Map[T any, R any](src []T, f func(T) R) []R {
	var dst []R
	for _, v := range src {
		dst = append(dst, f(v))
	}
	return dst
}

// Reduce 基于自定义函数聚合切片
func Reduce[T any, R any](src []T, f func(R, T) R) R {
	var init R
	for _, v := range src {
		init = f(init, v)
	}
	return init
}
