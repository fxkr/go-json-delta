package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffIntSuite struct {
}

func (s *DiffIntSuite) Test0(c *C) {
	left := 0
	right := 0
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffIntSuite) Test1(c *C) {
	left := 1
	right := 1
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffIntSuite) TestMinus1(c *C) {
	left := -1
	right := -1
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffIntSuite) Test0to1(c *C) {
	left := 0
	right := 1
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 1},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffIntSuite) Test1to0(c *C) {
	left := 1
	right := 0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffIntSuite) TestIntToNil(c *C) {
	left := 1
	var right interface{} = nil
	expected := []interface{}{
		[]interface{}{[]interface{}{}, nil},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
