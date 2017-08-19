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
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) Test1point1(c *C) {
	left := 1.1
	right := 1.1
	expected := []interface{}{}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) TestMinus1point1(c *C) {
	left := -1.1
	right := -1.1
	expected := []interface{}{}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
func (s *DiffFloatSuite) Test0point0to1point1(c *C) {
	left := 0.0
	right := 1.1
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 1.1},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) Test1point1to0point0(c *C) {
	left := 1.1
	right := 0.0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0.0},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) TestFloatToNil(c *C) {
	left := 0.0
	var right interface{} = nil
	expected := []interface{}{
		[]interface{}{[]interface{}{}, nil},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
