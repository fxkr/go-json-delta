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
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffStringSuite) TestNotEmpty(c *C) {
	left := "test"
	right := "test"
	expected := []interface{}{}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffStringSuite) TestEmptyToNotEmpty(c *C) {
	left := ""
	right := "test"
	expected := []interface{}{
		[]interface{}{[]interface{}{}, "test"},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffStringSuite) TestNotEmptyToEmpty(c *C) {
	left := "test"
	right := ""
	expected := []interface{}{
		[]interface{}{[]interface{}{}, ""},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffStringSuite) TestChange(c *C) {
	left := "foo"
	right := "bar"
	expected := []interface{}{
		[]interface{}{[]interface{}{}, "bar"},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffFloatSuite) TestStringToNil(c *C) {
	left := ""
	var right interface{} = nil
	expected := []interface{}{
		[]interface{}{[]interface{}{}, nil},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
