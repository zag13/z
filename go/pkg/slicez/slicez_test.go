package slicez

import "testing"

func TestSliceColumn(t *testing.T) {
	data := []map[string]interface{}{
		{
			"id":   1,
			"name": "张三",
		},
		{
			"id":   2,
			"name": "李四",
		},
		{
			"id":   3,
			"name": "王五",
		},
	}

	t.Log(SliceColumn(data, "id"))
	t.Log(SliceColumn(data, "name"))
}
