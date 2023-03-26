package slicez

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		src    []int
		f      func(int) bool
		expect []int
	}{
		{
			src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			f: func(i int) bool {
				return i%2 == 0
			},
			expect: []int{2, 4, 6, 8},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.src), func(t *testing.T) {
			assert.Equal(t, c.expect, Filter(c.src, c.f))
		})
	}
}
