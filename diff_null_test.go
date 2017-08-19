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
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) TestNullTo1(c *C) {
	var left interface{} = nil
	right := 1
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 1},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
