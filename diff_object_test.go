package jsondelta

import (
	. "gopkg.in/check.v1"
	"reflect"
)

type DiffObjectSuite struct {
}

func (s *DiffObjectSuite) TestEmpty(c *C) {
	left := map[string]interface{}{}
	right := map[string]interface{}{}
	expected := []interface{}{}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectToZero(c *C) {
	left := map[string]interface{}{}
	right := 0
	expected := []interface{}{
		[]interface{}{[]interface{}{}, 0},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectNewKey(c *C) {
	left := map[string]interface{}{}
	right := map[string]interface{}{"key1": "value1", "key2": "value2"}
	expectedItem1 := 	[]interface{}{[]interface{}{"key1"}, "value1"}
	expectedItem2 := []interface{}{[]interface{}{"key2"}, "value2"}
	expected1 := []interface{}{expectedItem1, expectedItem2}
	expected2:= []interface{}{expectedItem2, expectedItem1}
	obtained, err := diff(left, right) // Order of patches is not stable
	c.Assert(err, IsNil)
	equals1 := reflect.DeepEqual(obtained, expected1);
	equals2 := reflect.DeepEqual(obtained, expected2);
	c.Assert(equals1||equals2, Equals, true)
}

func (s *DiffObjectSuite) TestObjectRemovedKey(c *C) {
	left := map[string]interface{}{"key1": "value1", "key2": "value2"}
	right := map[string]interface{}{"key1": "value1"}
	expected := []interface{}{
		[]interface{}{[]interface{}{"key2"}},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectChangedValue(c *C) {
	left := map[string]interface{}{"key1": "value1"}
	right := map[string]interface{}{"key1": "value2"}
	expected := []interface{}{
		[]interface{}{[]interface{}{"key1"}, "value2"},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectObjectObjectValueChanged(c *C) {
	left := map[string]interface{}{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value1"}}}
	right := map[string]interface{}{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value2"}}}
	expected := []interface{}{
		[]interface{}{[]interface{}{"key1", "key2", "key3"}, "value2"},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}

func (s *DiffObjectSuite) TestObjectObjectObjectValueDeleted(c *C) {
	left := map[string]interface{}{"key1": map[string]interface{}{"key2": map[string]interface{}{"key3": "value1"}}}
	right := map[string]interface{}{"key1": map[string]interface{}{"key2": map[string]interface{}{}}}
	expected := []interface{}{
		[]interface{}{[]interface{}{"key1", "key2", "key3"}},
	}
	obtained, err := diff(left, right)
	c.Assert(err, IsNil)
	c.Assert(obtained, DeepEquals, expected)
}