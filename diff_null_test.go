package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffNullSuite struct {
}

func (s *DiffNullSuite) TestNull(c *C) {
	var left interface{} = nil
	var right interface{} = nil
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
