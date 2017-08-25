package jsondelta

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type DiffCustomSuite struct {
}

type customType struct {
	String string "json:`foo`"
}

type customTypeWithMarshal struct {
	String string "json:`foo`"
}

func (c *customTypeWithMarshal) MarshalJSON()([]byte, error) {
	return json.Marshal(map[string]interface{}{"bar": c.String})
}

func (s *DiffCustomSuite) TestCustomType(c *C) {
	left := customType{"a"}
	right := customType{"b"}
	expected := []interface{}{
		[]interface{}{[]interface{}{"String"}, "b"},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffCustomSuite) TestCustomPointerType(c *C) {
	left := &customType{"a"}
	right := &customType{"b"}
	expected := []interface{}{
		[]interface{}{[]interface{}{"String"}, "b"},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffCustomSuite) TestCustomPointerTypeWithMarshal(c *C) {
	left := &customTypeWithMarshal{"a"}
	right := &customTypeWithMarshal{"b"}
	expected := []interface{}{
		[]interface{}{[]interface{}{"bar"}, "b"},
	}
	obtained, err := Diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}