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
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffBoolSuite) TestFalse(c *C) {
	left := false
	right := false
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
