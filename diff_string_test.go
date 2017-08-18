package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffStringSuite struct {
}

func (s *DiffStringSuite) TestEmpty(c *C) {
	left := ""
	right := ""
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffStringSuite) TestNotEmpty(c *C) {
	left := "test"
	right := "test"
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
