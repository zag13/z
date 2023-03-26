package slicez

import "testing"

func TestColumn(t *testing.T) {
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

	t.Log(Column(data, "id"))
	t.Log(Column(data, "name"))
}
