package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffObjectSuite struct {
}

func (s *DiffObjectSuite) TestEmpty(c *C) {
	left := map[string]interface{}{}
	right := map[string]interface{}{}
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectToZero(c *C) {
	left := map[string]interface{}{}
	right := 0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
