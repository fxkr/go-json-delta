package jsondelta

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	_ = Suite(&DiffArraySuite{})
	_ = Suite(&DiffBoolSuite{})
	_ = Suite(&DiffFloatSuite{})
	_ = Suite(&DiffIntSuite{})
	_ = Suite(&DiffNullSuite{})
	_ = Suite(&DiffObjectSuite{})
	_ = Suite(&DiffStringSuite{})
	TestingT(t)
}
