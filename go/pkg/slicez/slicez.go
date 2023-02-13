package slicez

// SliceColumn 从切片中取出某个字段的值
func SliceColumn(input []map[string]interface{}, fieldName string) (result []interface{}) {
	for _, v := range input {
		result = append(result, v[fieldName])
	}
	return result
}
