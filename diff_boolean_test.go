package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffBoolSuite struct {
}

func (s *DiffBoolSuite) TestTrue(c *C) {
	left := true
	right := true
	expected := []interface{}{}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffBoolSuite) TestFalse(c *C) {
	left := false
	right := false
	expected := []interface{}{}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffBoolSuite) TestFalseToTrue(c *C) {
	left := false
	right := true
	expected := []interface{}{
		[]interface{}{[]interface{}{}, true},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffBoolSuite) TestTrueToFalse(c *C) {
	left := true
	right := false
	expected := []interface{}{
		[]interface{}{[]interface{}{}, false},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffBoolSuite) TestTrueToZero(c *C) {
	left := false
	right := 0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
