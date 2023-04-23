package assert_test

import (
	"errors"
	"testing"

	"github.com/fabioelizandro/manitable/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AssertSuite struct {
	*require.Assertions
	suite.Suite
}

func TestAssertSuite(t *testing.T) {
	s := new(AssertSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *AssertSuite) TestNoErr() {
	assert.NoErr(nil)

	f := func() {
		assert.NoErr(errors.New("foo"))
	}

	s.Panics(f)
}

func (s *AssertSuite) TestTrue() {
	assert.True(true, "s")

	f := func() {
		assert.True(false, "t")
	}

	s.Panics(f)
}

func (s *AssertSuite) TestFalse() {
	assert.False(false, "s")

	f := func() {
		assert.False(true, "t")
	}

	s.Panics(f)
}

func (s *AssertSuite) TestUnreachable() {
	f := func() {
		assert.Unreachable("t")
	}

	s.Panics(f)
}
