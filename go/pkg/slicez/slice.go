package slicez

// Index 返回元素在切片中的索引，如果不存在则返回-1
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// Remove 删除切片中的元素
func Remove[T comparable](s []T, x T) []T {
	i := Index(s, x)
	if i == -1 {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// RemoveDuplicate 删除切片中的重复元素
func RemoveDuplicate[T comparable](s []T) []T {
	var result []T
	m := make(map[T]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Intersect 求两个切片的交集
func Intersect[T comparable](s1, s2 []T) []T {
	var result []T
	m := make(map[T]bool)
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// Column 从二维切片中取出某个字段所有的值
func Column(input []map[string]interface{}, fieldName string) (result []interface{}) {
	for _, v := range input {
		result = append(result, v[fieldName])
	}
	return result
}
