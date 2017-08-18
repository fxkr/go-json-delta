package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffFloatSuite struct {
}

func (s *DiffFloatSuite) Test0point0(c *C) {
	left := 0.0
	right := 0.0
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) Test1point1(c *C) {
	left := 1.1
	right := 1.1
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) TestMinus1point1(c *C) {
	left := -1.1
	right := -1.1
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
