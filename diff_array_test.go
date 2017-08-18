package jsondelta

import (
	. "gopkg.in/check.v1"
)

type DiffArraySuite struct {
}

func (s *DiffArraySuite) TestEmpty(c *C) {
	left := []interface{}{}
	right := []interface{}{}
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffArraySuite) TestNested(c *C) {
	left := []interface{}{[]interface{}{}}
	right := []interface{}{[]interface{}{}}
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffArraySuite) TestNestedTwice(c *C) {
	left := []interface{}{[]interface{}{[]interface{}{}}}
	right := []interface{}{[]interface{}{[]interface{}{}}}
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestArrayToZero(c *C) {
	left := map[string]interface{}{}
	right := 0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}
